# Chicago Business Intelligence Project

## Programs/Microservices Deployed
1. Google Cloud Platform - houses the microservices
    a. Postgres Admin
    b. Go Services
2. Google Docker Registry
3. Google Cloud Registry
4. GitHub Repository
5. Postgres 


## How to Install and Run the Project
1. Clone the repository from GitHub: https://github.com/JoleneMaxa/chicago_business_intelligence.git
2. Pull it down to VS Code and make its own directory.
3. Install Google CLI.
4. Log into your Google Cloud Console.
5. Create a new project.
6. Create a new Postgres instance of the database, along with new user, and a new database.
7. Enable/create the GitHub repository push trigger on the Google Cloud build API.
8. Enable the Cloud Run API, which will set up the microservice containers.
9. Ensure proper connections/permissions in GCP.
10. Put Postgres connection name in the main.go script. Also Amend the cloudbuild.yaml file with the new Google Cloud Project ID.
11. Push script changes in VS Code to GitHub. 
12. Trigger will run automatically.
13. Ensure the Build completes in the Google Build API. This gets the Docker images, and builds them within the Google Cloud Repository for the 2 microservices we want to deploy.
14. Go to Cloud Run to confirm they are both running.
15. Then go check out the logs for each. Make sure the output is what you want.
16. Grab the links from each. One of them will be Postgre Admin to check out the data you pulled is in the database.

