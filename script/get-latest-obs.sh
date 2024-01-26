#!/bin/sh

main() {
        path=${1?specify where you want the deb}

        download_artifact >/tmp/obs.zip
        rm -rf /tmp/obs
        unzip -od /tmp/obs /tmp/obs.zip
        mv /tmp/obs/*.deb "$path"
        ls -al "$path"
}

download_artifact() {
        repo=obsproject/obs-studio
        id=$(gh api /repos/$repo/actions/artifacts --jq '
                [
                        .artifacts[]
                        | select(
                                .workflow_run.head_branch == "master" and
                                (.name|contains("ubuntu")) and
                                (.name|contains("dbgsym")|not)
                        )
                ]
                | sort_by(.created_at)
                | reverse
                | first
                | .id
        ')
        gh api "/repos/$repo/actions/artifacts/$id/zip"
}

set -eu
main "$@"
