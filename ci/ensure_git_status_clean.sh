if [ -z "$(git status --porcelain)" ]; then
  exit 0
else
  echo "Uncommit changes present:"
  git status
  exit 1
fi
