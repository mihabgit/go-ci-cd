# Go CI/CD Example

This is a simple Golang project demonstrating a basic Continuous Integration and Continuous Delivery (CI/CD) pipeline using GitHub Actions.

## Overview

This project consists of a simple Go program that greets a given name and includes a basic unit test. The GitHub Actions workflow automates the following steps upon code changes:

- **Building** the Go project.
- **Testing** the Go project.
- A basic **Delivery** step that prints a message after successful build and test (representing a potential deployment stage).

## Getting Started

1.  **Clone the repository:**
    ```bash
    git clone [https://github.com/your-username/go-ci-cd.git](https://github.com/your-username/go-ci-cd.git)
    cd go-ci-cd
    ```
    (Replace `https://github.com/your-username/go-ci-cd.git` with your repository URL)

2.  **Explore the Go code:**
    - `main.go`: Contains the main application logic with a `Greet` function.
    - `main_test.go`: Contains a unit test for the `Greet` function.

## CI/CD Pipeline

The CI/CD pipeline is defined in the `.github/workflows/go.yaml` file.

### Workflow Triggers

The workflow is triggered on:

-   `push` events to the `main` branch.
-   `pull_request` events targeting the `main` branch.

### Jobs

The workflow consists of two jobs:

1.  **`build`**:
    -   Checks out the code.
    -   Sets up the Go environment.
    -   Downloads Go module dependencies.
    -   Builds the Go project.
    -   Runs the unit tests.

2.  **`deliver`**:
    -   This job only runs if the `build` job succeeds (`needs: build`).
    -   Prints a simple message indicating that the build and tests were successful. In a real-world scenario, this job could be extended to deploy the application.

## How to See the CI/CD in Action

1.  **Make changes to the Go code** (e.g., modify `main.go` or `main_test.go`).
2.  **Commit and push your changes** to the `main` branch of your GitHub repository:
    ```bash
    git add .
    git commit -m "Make some changes"
    git push origin main
    ```
3.  **Create a new Pull Request** targeting the `main` branch.

You can then navigate to the "Actions" tab of your GitHub repository to see the CI/CD pipeline running. You'll see the progress of the `build` and `deliver` jobs, and whether they succeed or fail.

## Next Steps and Further Learning

This is a very basic CI/CD example. To further enhance your learning, you can explore:

-   **Adding more sophisticated tests:** Integration tests, linters, security scans.
-   **Automating releases:** Tagging releases, building binaries.
-   **Containerization with Docker:** Building and pushing Docker images in your pipeline.
-   **Deployment to different environments:** Configuring deployment to staging and production environments using cloud platforms or other deployment tools.
-   **Managing environment variables and secrets** securely in your GitHub Actions workflows.
-   **Exploring different workflow triggers** and conditions.

This project provides a hands-on starting point for understanding and implementing CI/CD principles with Golang and GitHub Actions. Happy coding!
