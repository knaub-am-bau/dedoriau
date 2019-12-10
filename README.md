# dedoriau
Decentralized Docker Registry Authentication

This project aims to have a authentication outside the Docker environment.

In this first approach we use the Iota Tangle.

The Provider of the container has to send the ID and the Docker hash to the Tangle.
The Client requests the container hash with the Docker ID from the Tangle.
With this authenticated docker hash, the right container can be downloaded.

![functional plan](/dedoriau.jpg)
