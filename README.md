# gem

Out of paper? No problem! Jot down your ideas right here in the terminal with `gem`.

## Features
- **Add Ideas**: Quickly add your ideas to the gem storage.
- **List Ideas**: View your stored ideas easily.
- **Remove Ideas**: Delete ideas that you no longer need.

## Usage

### Add an Idea
To add an idea to the gem storage, use the `-add` flag:
```
gem -add "<your idea>"
```
Example:
```
gem -add "Build a new feature for my app"
```

### List Ideas
To list all the ideas stored in gem, use the `-ls` and `-all` flag:
```
gem -ls -all
```
You can also specify the item :
```
gem -ls <number of item>
```

## Usage 

### List item number 5
```
gem -ls 5
```
### List all items
```
gem -ls -all
```

### Remove an Idea
To remove an idea from gem storage, use the `-rm` flag followed by the number of the idea from the list:
```
gem -rm <number of idea from the list>
```
Example:
```
gem -rm 5
```

## About
`gem` is a simple and fast way to manage your ideas directly from the terminal. This is my first Go project, and I'm excited to share it with you! :D


