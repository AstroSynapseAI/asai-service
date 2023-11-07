# ASAI App Service 

Core backend service providing application and AI engine API. App service supports rest API for authentication, AI engine prompting, and managing the application data and settings. It also supports web socket connection for streaming AI engine results.

The AI engine is currently powered by openai models, with plans to implement Ollama open source models support soon.

## Development setup

Clone the repository and install dependencies:

```bash
git clone https://github.com/astrosynapseai/app-service
cd app-service
go get
go mod tidy
```

To authenticate for openai APIs and APIs the agents use, you need to rename `keys.template.yaml` to `keys.yaml` and input the required API keys.

```yaml
open_api_key:    ""
serpapi_api_key: ""
discord_api_key: ""
```

- For openpi API key visit: https://platform.openai.com/
- For serpapi API key visit: https://serpapi.com/
- For discord API key visit: https://discordgsm.com/guide/how-to-get-a-discord-bot-token

### Environments

For local development, the repository comes with a docker-compose.yaml and a Docker file for running the app service with hot reload support [using go air package](https://github.com/cosmtrek/air). To setup the environment make sure you have docker and docker-compose installed on your local machine, then from the root run: 

```bash
docker compose build
docker compose up
```

> Main development branch is `dev` branch, all changes merged into `main` branch will automatically be deployed to https://asai.astrosynapse.ai

## Web Client

Vue based web client is located in `./web` folder. To install dependacies run:

```bash
cd web
npm install
```

To run with docker the web client needs to be built "for production" with `npm run build`, the build will be saved in `./web/static` folder.

