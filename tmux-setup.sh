#!/bin/bash

# Define the session name
SESSION_NAME="downloads-cli"

# Start a new tmux session with the specified name or attach if it already exists
tmux has-session -t $SESSION_NAME 2>/dev/null
if [ $? != 0 ]; then
  # Create a new session
  tmux new-session -d -s $SESSION_NAME -c ~/code/downloads-cli

  # Create a window for nvim
  tmux rename-window -t $SESSION_NAME:0 "Window"
  tmux send-keys -t $SESSION_NAME:0 "Window" C-m

  # Split the pane
  tmux split-window -h -t $SESSION_NAME:0
  tmux resize-pane -t $SESSION_NAME:0.1 -x 10%
  tmux send-keys -t $SESSION_NAME:0.1 "gow -c run ." C-m

  # Create an empty window
  tmux new-window -t $SESSION_NAME -n "Window"
fi

# Select the "nvim" window before attaching
tmux select-window -t $SESSION_NAME:0

# Attach to the session
tmux attach-session -t $SESSION_NAME
