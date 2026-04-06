#!/bin/bash


clone(){
	echo "Code Cloning..."
	git clone https://github.com/nasirhayat028/resume.git || return 1
	echo "Code Successfully cloned"
}

<<simple
install_requirements() {
	echo "Installing dependencies..."
    	sudo apt-get update && sudo apt-get install -y docker.io nginx docker-compose || {
        echo "Failed to install dependencies."
        return 1
    }
}
simple


required_restarts() {
    echo "Performing required restarts..."
    sudo chown "$USER" /var/run/docker.sock || {
        echo "Failed to change ownership of docker.sock."
        return 1
    }

    # Uncomment the following lines if needed:
    # sudo systemctl enable docker
    # sudo systemctl enable nginx
    # sudo systemctl restart docker
}


deploy() {
    echo "Building and deploying the resume app..."
    docker build -t resume-app . && docker-compose up -d || {
      	echo "Failed to build and deploy the app."
        return 1
    }
}


# Main deployment script
echo "********** DEPLOYMENT STARTED *********"

# Clone the code
if ! clone; then
    cd resume || exit 1
fi
<< comment
# Install dependencies
if ! install_requirements; then
    exit 1
fi
comment

# Perform required restarts
if ! required_restarts; then
    exit 1
fi

# Deploy the app
if ! deploy; then
    echo "Deployment failed. Mailing the admin..."
    # Add your sendmail or notification logic here
    exit 1
fi

echo "********** DEPLOYMENT DONE *********"
