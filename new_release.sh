#!/bin/bash

bump_type=${1:-patch}

latest_tag=$(git tag --sort=-v:refname | head -n1)

if [ -z "$latest_tag" ]; then
    major=0
    minor=0
    patch=0
else
    IFS='.' read -r major minor patch <<<"${latest_tag#v}"
fi

case $bump_type in
    major)
        major=$((major + 1))
        minor=0
        patch=0
        ;;
    minor)
        minor=$((minor + 1))
        patch=0
        ;;
    patch)
        patch=$((patch + 1))
        ;;
    *)
        echo "Usage: $0 [major|minor|patch]"
        exit 1
        ;;
esac

new_tag="v${major}.${minor}.${patch}"

while git rev-parse "$new_tag" >/dev/null 2>&1; do
    patch=$((patch + 1))
    new_tag="v${major}.${minor}.${patch}"
done

echo "Last tag: $latest_tag"
echo "New tag: $new_tag"

git tag "$new_tag"
git push origin "$new_tag"
