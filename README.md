# Hello, Gemini using Vertex AI API.

Sample app showing how to access Gemini using the Vertex AI API.


## Prerequisites

1.  Download and install Go, see https://go.dev/. To determine the version of Go
    that is available on your path run:

    ```sh
    go version
    ```

1.  [Optional] Install `gcloud` CLI, see https://cloud.google.com/sdk/docs/install.


## Prepare sample project

1.  Clone this repo.

    ```shell
    git clone git@github.com:fredsa/hello-gemini-vertex-ai-api.git
    ```

1.  Change into the project directory.

    ```shell
    cd hello-gemini-vertex-ai-api
    ```

1.  Install the
    [`cloud.google.com/go/vertexai/genai`](https://pkg.go.dev/cloud.google.com/go/vertexai/genai)
    package.

    ```shell
    go get cloud.google.com/go/vertexai/genai
    ```


## Authorize the app

1.  To use
    [Application Default Credentials](https://cloud.google.com/docs/authentication/application-default-credentials)
    to authorize the app, run the following `gcloud` command:

    ```shell
    gcloud auth application-default login
    ```

    Alternatively, you can authorize the app using an API key by modifying `main.go` to use
    [`option.WithAPIKey`](https://pkg.go.dev/google.golang.org/api/option#WithAPIKey) when
    creating the client:

    ```go
    // New client, authorize using an API KEY.
    option := option.WithAPIKey("YOUR_API_KEY")
    client, err := genai.NewClient(ctx, project, location, option)
    ```

## Run the sample

1. Compile and run.

    ```shell
    go run main.go
    ```

    The output should look something like this
    ```none
    >> Hello, who are you?
    I am a large language model, trained by Google.
    ```