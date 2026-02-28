#!/bin/sh
#
# Compares two obs-websocket protocol versions and outputs a human-readable
# summary of what changed. Used by the bump automation to generate PR descriptions.
#
# Usage: ./script/protocol-diff.sh <old_version> <new_version>

set -eu

old_version="$1"
new_version="$2"

tmpdir=$(mktemp -d)
trap 'rm -rf "$tmpdir"' EXIT

# Fetch protocol JSON from GitHub Contents API
fetch() {
        gh api "/repos/obsproject/obs-websocket/contents/docs/generated/protocol.json?ref=$1" \
                --jq .content | base64 -d > "$2"
}
fetch "$old_version" "$tmpdir/old.json"
fetch "$new_version" "$tmpdir/new.json"

# Extract sorted unique values from JSON via jq
extract() {
        jq -r "$2" "$1" | sort -u
}

# Diff two sorted lists: added (comm -13) or removed (comm -23)
diff_lists() {
        comm "$1" "$2" "$3"
}

section() {
        header="$1"
        items="$2"
        if [ -z "$items" ]; then return; fi
        echo "### $header"
        echo
        echo "$items" | while IFS= read -r item; do
                echo "- $item"
        done
        echo
}

old="$tmpdir/old.json"
new="$tmpdir/new.json"

# --- Request categories ---
extract "$old" '[.requests[].category] | unique | .[]' > "$tmpdir/old_req_cat"
extract "$new" '[.requests[].category] | unique | .[]' > "$tmpdir/new_req_cat"
added_req_cat=$(diff_lists -13 "$tmpdir/old_req_cat" "$tmpdir/new_req_cat")
removed_req_cat=$(diff_lists -23 "$tmpdir/old_req_cat" "$tmpdir/new_req_cat")

# --- Requests ---
extract "$old" '.requests[] | "\(.category)/\(.requestType)"' > "$tmpdir/old_req"
extract "$new" '.requests[] | "\(.category)/\(.requestType)"' > "$tmpdir/new_req"
added_req=$(diff_lists -13 "$tmpdir/old_req" "$tmpdir/new_req")
removed_req=$(diff_lists -23 "$tmpdir/old_req" "$tmpdir/new_req")

# --- Event categories ---
extract "$old" '[.events[].category] | unique | .[]' > "$tmpdir/old_evt_cat"
extract "$new" '[.events[].category] | unique | .[]' > "$tmpdir/new_evt_cat"
added_evt_cat=$(diff_lists -13 "$tmpdir/old_evt_cat" "$tmpdir/new_evt_cat")
removed_evt_cat=$(diff_lists -23 "$tmpdir/old_evt_cat" "$tmpdir/new_evt_cat")

# --- Events ---
extract "$old" '.events[] | "\(.category)/\(.eventType)"' > "$tmpdir/old_evt"
extract "$new" '.events[] | "\(.category)/\(.eventType)"' > "$tmpdir/new_evt"
added_evt=$(diff_lists -13 "$tmpdir/old_evt" "$tmpdir/new_evt")
removed_evt=$(diff_lists -23 "$tmpdir/old_evt" "$tmpdir/new_evt")

# --- Enum types ---
extract "$old" '.enums[].enumType' > "$tmpdir/old_enum"
extract "$new" '.enums[].enumType' > "$tmpdir/new_enum"
added_enum=$(diff_lists -13 "$tmpdir/old_enum" "$tmpdir/new_enum")
removed_enum=$(diff_lists -23 "$tmpdir/old_enum" "$tmpdir/new_enum")

# --- Enum values ---
extract "$old" '.enums[] | .enumType as $t | .enumIdentifiers[] | "\($t)/\(.enumIdentifier)"' > "$tmpdir/old_enumv"
extract "$new" '.enums[] | .enumType as $t | .enumIdentifiers[] | "\($t)/\(.enumIdentifier)"' > "$tmpdir/new_enumv"
added_enumv=$(diff_lists -13 "$tmpdir/old_enumv" "$tmpdir/new_enumv")
removed_enumv=$(diff_lists -23 "$tmpdir/old_enumv" "$tmpdir/new_enumv")

# --- Deprecations ---
extract "$old" '(.requests[] | select(.deprecated) | .requestType), (.events[] | select(.deprecated) | .eventType)' > "$tmpdir/old_dep"
extract "$new" '(.requests[] | select(.deprecated) | .requestType), (.events[] | select(.deprecated) | .eventType)' > "$tmpdir/new_dep"
newly_dep=$(diff_lists -13 "$tmpdir/old_dep" "$tmpdir/new_dep")

# --- Output ---
has_changes=false
for val in "$added_req_cat" "$removed_req_cat" "$added_req" "$removed_req" \
           "$added_evt_cat" "$removed_evt_cat" "$added_evt" "$removed_evt" \
           "$added_enum" "$removed_enum" "$added_enumv" "$removed_enumv" "$newly_dep"; do
        if [ -n "$val" ]; then has_changes=true; break; fi
done

if [ "$has_changes" = false ]; then
        echo "No semantic changes between \`$old_version\` and \`$new_version\`."
        exit 0
fi

echo "## Protocol diff: \`$old_version\` → \`$new_version\`"
echo

section "New request categories" "$added_req_cat"
section "Removed request categories" "$removed_req_cat"
section "New requests" "$added_req"
section "Removed requests" "$removed_req"
section "New event categories" "$added_evt_cat"
section "Removed event categories" "$removed_evt_cat"
section "New events" "$added_evt"
section "Removed events" "$removed_evt"
section "New enum types" "$added_enum"
section "Removed enum types" "$removed_enum"
section "New enum values" "$added_enumv"
section "Removed enum values" "$removed_enumv"
section "Newly deprecated" "$newly_dep"
