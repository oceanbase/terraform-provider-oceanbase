#!/usr/bin/env bash
set -euo pipefail

demo_dir="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
repo_dir="$(cd "$demo_dir/../.." && pwd)"

mkdir -p "$demo_dir/.provider"
(cd "$repo_dir" && go build -o "$demo_dir/.provider/terraform-provider-oceanbase" ./main.go)

cat > "$demo_dir/.terraformrc" <<EOF
provider_installation {
  dev_overrides {
    "oceanbase.github.io/oceanbase/oceanbase" = "$demo_dir/.provider"
  }
  direct {}
}
EOF

printf 'Provider built. In this shell run:\nexport TF_CLI_CONFIG_FILE="%s/.terraformrc"\n' "$demo_dir"
