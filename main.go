package main

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"os"

	"cloud.google.com/go/vertexai/genai"
)

const project = "your-project-id" // Cloud project ID.
const location = "us-central1"    // Cloud project location.
const modelName = "gemini-pro"    // OR "gemini-pro-vision".

func main() {
	ctx := context.Background()

	// New client, using Application Default Context (ADC).
	client, err := genai.NewClient(ctx, project, location)
	if err != nil {
		log.Fatalf("Error creating client: %v\n", err)
	}
	defer client.Close()

	// Start conversation with Gemini.
	converse(ctx, client.GenerativeModel(modelName))
}

func converse(ctx context.Context, model *genai.GenerativeModel) {
	// Get user input from stdin.
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("\n>> ")

		// Read user prompt.
		prompt, err := reader.ReadString('\n')
		if err != nil {
			log.Fatalf("Error reading input: %v\n", err)
		}

		// Call the Vertex AI API.
		resp, err := model.GenerateContent(ctx, genai.Text(prompt))
		if err != nil {
			log.Fatalf("Error generating content: %v\n", err)
		}

		// Display the response.
		for _, part := range resp.Candidates[0].Content.Parts {
			fmt.Printf("%v\n", part)
		}
	}
}
