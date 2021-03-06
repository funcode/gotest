# /etc/launchd.conf sets the inital PATH

# disable git status check
export SCM_CHECK=false
# git config --global --add bash-it.hide-status 0

# Path to the bash it configuration
export BASH_IT="/Users/pivotal/.bash_it"

# Lock and Load a custom theme file
export BASH_IT_THEME="demula"

# Load Bash It
source $BASH_IT/bash_it.sh
test -e "${HOME}/.iterm2_shell_integration.bash" && source "${HOME}/.iterm2_shell_integration.bash"

# export GOROOT=/usr/local/go
export GOROOT=/usr/local/Cellar/go/1.9.2/libexec/
export GOPATH=/Users/pivotal/workspace
export PATH=$GOROOT/bin:$GOPATH/bin:$PATH

# function gfm()
# {
#   f=$1
#   echo "Processing $f..."
# }

# export gfm
# alias gfm="vim -c 'LivedownPreview' "
alias vf="pushd ~; rm .vim; ln -s ~/workspace/vimfiles .vim; rm .vimrc; ln -s ~/workspace/vimfiles/vimrc .vimrc;      rm ~/.vimrc.local; ln -s ~/workspace/go/.vimrc.local.vf .vimrc.local; popd"
alias vc="pushd ~; rm .vim; ln -s ~/workspace/vim-config .vim; rm .vimrc; ln -s ~/workspace/vim-config/vimrc .vimrc;  rm ~/.vimrc.local; ln -s ~/workspace/go/.vimrc.local.vc .vimrc.local; popd"
alias vl="ls -l ~/.vim;ls -l ~/.vimrc*"
alias viml="vim ~/.vimrc.local"
alias vimrc="vim ~/.vimrc"
alias vb="vim ~/.bash_profile"


