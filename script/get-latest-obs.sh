#!/bin/sh

main() {
        path=${1?specify where you want the deb}

        download_artifact >/tmp/obs.zip
        rm -rf /tmp/obs
        unzip -od /tmp/obs /tmp/obs.zip

        # the names are different in their sha256 file
        expected=$(cut -d' ' -f1 /tmp/obs/*.sha256 | awk '{print $1}')
        mv /tmp/obs/*.deb "$path"
        echo "$expected $path" | sha256sum --check
        ls -al "$path"
}

download_artifact() {
        repo=obsproject/obs-studio
        artifacts_url=$(
                gh api -XGET /repos/$repo/actions/runs \
                        -fbranch=master \
                        -fstatus=completed \
                        -fexclude_pull_requests=true \
                        -fevent=push \
                        --jq '
                                [
                                        .workflow_runs[]
                                        | select(.path==".github/workflows/push.yaml")
                                ]
                                | sort_by(.created_at)
                                | reverse
                                | first
                                | .artifacts_url
                        '
        )
        archive_download_url=$(
                gh api "$artifacts_url" --jq '
                        .artifacts[]
                        |= select(.name | test("ubuntu-2[4-9].[0-9]{2}-x86_64-[0-9a-f]+$"))
                        | .artifacts
                        | first
                        | .archive_download_url
                '
        )
        gh api "$archive_download_url"
}

set -eux
main "$@"
