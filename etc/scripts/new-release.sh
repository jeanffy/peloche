if [ -z "$1" ]; then
  echo "Tag name vX.Y.Z missing"
  exit 1
fi

TAG_NAME=$1
echo "New release '$TAG_NAME'"

git tag $TAG_NAME -a -m "$TAG_NAME"
git push tag $TAG_NAME
