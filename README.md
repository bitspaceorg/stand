<div align=center><h1> STAND</h1>

  <img src="https://img.shields.io/badge/:bitspace x fosshack-%23121011?style=for-the-badge&logoColor=%23ffffff&color=%23000000">
    </div>

## Table of Contents

- [About](#about)
- [Installation](#installation)
- [Usage](#usage)
- [Features](#features)
- [Configuration](#configuration)
- [Techstack](#techstack)
- [Contributing](#contributing)
- [License](#license)

## About

STAND (Scalable Technology for Advanced Network Deployment) is a self hosting open source project designed for on-premise servers.
It allows users to deploy applications directly from GitHub using a simple and intuitive web UI.
STAND aims to simplify the process of deploying and managing applications on local servers, providing a scalable and efficient solution.
All the data required to deploy and host the applications is stored locally, ensuring data privacy and security and most importantly providing the user with full control over their data.

## Installation

To install STAND, follow these steps:

1. Clone the repository:
   ```bash
   git clone https://github.com/bitspaceorg/stand-fosshack.git
   ```
2. Navigate to the project directory:
   ```bash
   cd stand-fosshack
   ```
3. Install the application:
   ```bash
   make install
   ```

## Usage

To use STAND, follow these steps:

1. Initialize the application:
   ```bash
   stand init
   ```
2. Start the application:
   ```bash
   stand start
   ```
3. Open your web browser and navigate to `http://localhost:3000`.
4. Download the GitHub app as per the instructions provided on the web UI.
5. Start hosting your projects.

## Features

- **Automated Configuration:** Automatically generates configuration files based on user input through the web UI, eliminating the need for manual edits.
- **Scalability:** Easily manage and scale your on-premise servers and applications.
- **Simple Web UI:** Intuitive and user-friendly interface for deploying applications directly from GitHub repositories.
- **Customizable Deployment:** Supports various programming languages and versions, with customizable build and run commands.
- **Efficient Logging:** Configurable log directories for easy access and management of application logs.
- **Environment Management:** Easily set and manage environment variables for your applications.
- **Quick Initialization:** Simple commands (`stand init` and `stand start`) to initialize and start the application, getting you up and running quickly.
- **Secure Hosting:** Provides options for enabling SSL to secure your deployments.

## Configuration

STAND generates the following configuration automatically based on user input from the web UI:

```yaml
project:
  name: your_project_name
  home: /path/to/your/project
  log: /path/to/log/directory/
requirements:
  language: your_language
  version: your_version
build:
  - name: install_dependencies
    cmd: your_install_command
run:
  - name: start_application
    cmd: your_start_command
env:
  - name: ENV_VARIABLE_NAME
    value: your_value
```

This configuration file is generated automatically and does not need to be manually edited. Users can specify project details, dependencies, build commands, run commands, and environment variables through the web UI, and STAND will handle the rest.

## Techstack

We used GO for its exceptional performance and efficiency, GO also provides great support for concurrency and parallelism which is essential for a project like STAND.

![Go](https://img.shields.io/badge/go-%2300ADD8.svg?style=for-the-badge&logo=go&logoColor=white)

The reason for using Next.js is its ability to provide a great user experience with its server-side rendering capabilities and its ability to provide a great developer experience with its hot reloading feature.

![Next JS](https://img.shields.io/badge/Next-black?style=for-the-badge&logo=next.js&logoColor=white)

Currently, STAND supports NodeJS applications, but we plan to add support for more languages and frameworks in the future.

![NodeJS](https://img.shields.io/badge/node.js-6DA55F?style=for-the-badge&logo=node.js&logoColor=white)

NPM is used to manage the dependencies of the node applications.

![NPM](https://img.shields.io/badge/NPM-%23CB3837.svg?style=for-the-badge&logo=npm&logoColor=white)

## Contributing

We welcome contributions from the community! If you would like to contribute to STAND, please follow the steps below:

1. Fork the repository
2. Clone the repository
3. Create a new branch
4. Make your changes
5. Commit your changes
6. Push your changes to your fork
7. Create a pull request explaining your changes

If possible try to crete an issue before making a pull request so that we can discuss the changes and make sure that they are in line with the project goals.
Note: Please make sure to follow the code of conduct and the contribution guidelines when contributing to STAND.

## License

STAND is licensed under the GNU GENERAL PUBLIC LICENSE . See the LICENSE file for more details.
