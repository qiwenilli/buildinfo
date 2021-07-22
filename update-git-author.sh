#!/bin/bash

git filter-branch --env-filter "
GIT_AUTHOR_NAME='qiwen'
GIT_AUTHOR_EMAIL='342145399@qq.com'
GIT_COMMITTER_NAME='qiwen'
GIT_COMMITTER_EMAIL='34214399@qq.com'
" -f



