# pb-go-app-template

This is a template repository designed to jumpstart Pocketbook application development using Go.

**Purpose**

This repository provides a pre-configured environment and basic structure for creating Pocketbook applications written in Go.  It aims to streamline the development process and get you started quickly.

**Key Features**

•   **Pocketbook Go Development:**  Provides a ready-to-use template for creating applications for Pocketbook devices using the Go programming language.
•   **Docker-Based Build Environment:**  Leverages a Docker image ([https://github.com/Skeeve/SDK_6.3.0](https://github.com/Skeeve/SDK_6.3.0)) to ensure a consistent and reproducible build environment.  This eliminates dependency conflicts and simplifies the build process.
•   **SDK Integration:** Utilizes the [https://github.com/Skeeve/SDK_6.3.0](https://github.com/Skeeve/SDK_6.3.0) Go library for interacting with the Pocketbook SDK.
•   **Makefile Automation:**  Includes a `Makefile` with pre-defined targets for common tasks such as:
    •   Building the application
    •   Automatically deploying the application to a connected Pocketbook device.
•   **Example GUI Application:**  Demonstrates a basic GUI application that:
    •   Displays text.
    •   Implements localization (i18n) by utilizing translations based on the system language setting of the Pocketbook device.
    •   Requests network connection permission.

**Getting Started**

1.  **Clone this repository:** Use it as a template to create your own repository.
2.  **Install Docker:** Make sure you have Docker installed on your system.
3.  **Connect your Pocketbook:** Ensure your Pocketbook device is connected to your computer via USB.
4.  **Explore the code:**  Examine the example application to understand the basic structure and how to use the SDK.
5.  **Build and Deploy:** Use the `make` commands in the `Makefile` to build and deploy your application to the connected Pocketbook.
6.  **Customize:** Modify the code to create your own Pocketbook application.

**License**

This repository is licensed under the MIT License.  You are free to copy, modify, and distribute it as you see fit.

**Share Your Creation!**

Once you've developed your awesome Pocketbook application, consider sharing it with the community! You can submit your application to the [pb-apps catalog](https://catinbeard.github.io/pb-apps/) to make it available to other Pocketbook users.