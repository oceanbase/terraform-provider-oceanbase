#!/usr/bin/env python3

import sys
import json

def get_referenced_models(spec, paths_to_keep):
    referenced_models = set()
    for path, methods in spec['paths'].items():
        if path not in paths_to_keep:
            continue
        for method in methods.values():
            if 'parameters' in method:
                for param in method['parameters']:
                    if 'schema' in param:
                        referenced_models.update(find_refs(param['schema'], spec))
            if 'responses' in method:
                for response in method['responses'].values():
                    if 'schema' in response:
                        referenced_models.update(find_refs(response['schema'], spec))
    return referenced_models

def find_refs(schema, spec, recursion_depth=100):
    refs = set()
    if recursion_depth == 0:
        return refs
    if isinstance(schema, dict):
        if '$ref' in schema:
            ref = schema['$ref'].split('/')[-1]
            refs.add(ref)
            # Recursively find refs in the referenced model
            if ref in spec.get('definitions', {}):
                refs.update(find_refs(spec['definitions'][ref], spec, recursion_depth=recursion_depth-1))
        for value in schema.values():
            refs.update(find_refs(value, spec, recursion_depth=recursion_depth-1))
    elif isinstance(schema, list):
        for item in schema:
            refs.update(find_refs(item, spec, recursion_depth=recursion_depth-1))
    return refs

def filter_swagger_spec(input_file, output_file, paths_to_keep):
    # Read the input file
    with open(input_file, 'r') as f:
        spec = json.load(f)

    # Filter paths
    filtered_paths = {path: spec['paths'][path] for path in paths_to_keep if path in spec['paths']}
    spec['paths'] = filtered_paths

    # Get referenced models
    referenced_models = get_referenced_models(spec, paths_to_keep)

    # Filter definitions (models)
    if 'definitions' in spec:
        filtered_definitions = {}
        for model in referenced_models:
            if model in spec['definitions']:
                filtered_definitions[model] = spec['definitions'][model]
        spec['definitions'] = filtered_definitions

    # Write the filtered spec to the output file
    with open(output_file, 'w') as f:
        json.dump(spec, f, indent=2)

if __name__ == "__main__":
    if len(sys.argv) < 4:
        print("Usage: python filter_swagger.py <input_file> <output_file> <path1> [<path2> ...]")
        sys.exit(1)

    input_file = sys.argv[1]
    output_file = sys.argv[2]
    paths_to_keep = sys.argv[3:]

    filter_swagger_spec(input_file, output_file, paths_to_keep)
    print(f"Filtered Swagger spec written to {output_file}")

