# GO-CLI

A command-line interface (CLI) application built with Go and Cobra for managing tasks.

## Features

- Add a task
- Get a task
- Update a task
- Delete a task

## Installation

1. Clone the repository:

    ```sh
    git clone https://github.com/samridhi-sahu/go-todo-cli.git
    ```

2. Navigate to the project directory:

    ```sh
    cd go-todo-cli
    ```

3. Install dependencies:

    ```sh
    go mod tidy
    ```

## Usage

### Add a Task
To add a task, use the `add` command with the `-n` flag:
```sh
go run main.go add "taskId" "taskName" "taskDesc" "taskStatus"
```
### Get a Task
To get a task, use the `get` command with the `-g` flag:
1. Get all tasks
```sh
go run main.go get
```
2. Get a specific task
```sh
go run main.go get "taskId"
```
3. Get multiple tasks
```sh
go run main.go get "taskId1" "taskId2" "taskId3"
```
### Update a Task
To update a task, use the `update` command with the `-u` flag:
```sh
go run main.go update "oldTaskId" "newTaskId" "newTaskName" "newTaskDesc" "newTaskStatus"
```
### Delete a Task
To delete a task, use the `delete` command with the `-d` flag:
```sh
go run main.go delete "taskId"
```

### Explanation

- **Project Title and Description**: Provides a brief overview of the project.
- **Features**: Lists the main features of the CLI application.
- **Installation**: Instructions to clone the repository and install dependencies.
- **Usage**: Examples of how to use each command.
- **Commands**: Lists the available commands and their purposes.
- **Flags**: Describes the flags used by each command.
- **License**: Mentions the project's license.

Feel free to customize this file to better suit your project's specifics.

## Contributing
Contributions are welcome! If you find any issues or have suggestions for improvements, please open an issue or submit a pull request.

## License
This project is licensed under the MIT License.
