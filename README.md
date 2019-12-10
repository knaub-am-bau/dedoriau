# dedoriau
Decentralized Docker Registry Authentication

This project aims to have a authentication outside the Docker environment.

In this first approach we use the Iota Tangle.

The Provider of the container has to send the ID and the Docker hash to the Tangle.
The Client requests the container hash with the Docker ID from the Tangle.
With this authenticated docker hash, the right container can be downloaded.

This Project is (for now) based on the code from the "go IOTA Workshop" guide
https://github.com/iota-community/go-iota-workshop.git

![functional plan](/dedoriau.jpg)
