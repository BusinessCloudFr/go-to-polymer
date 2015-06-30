# go-to-ploymer
This is a showcase using Golang App Engine, with Google Endpoint, and Polymer on the frontend.

Build for [Google-Polymer-paris Meetup](http://www.meetup.com/Google-Polymer-Paris/) by [Business-Cloud](http://www.business-cloud.fr) team.

## How to run the Project

### Prerequisite
To compile and run this project, you will need a [go appengine runtime](https://cloud.google.com/appengine/docs/go/gettingstarted/devenvironment) and [bower](http://bower.io/#install-bower).

### install dependencies

Javascript libraries are imported using bower, so from the root of your project
'''
cd ./frontend/static
bower update
'''

Import go dependencies
'''
goapp get github.com/GoogleCloudPlatform/go-endpoints/endpoints
'''

### Run the Demo
To run the demo, just serve the yamm files
'''
goapp serve ./dispatch.yaml frontend/app.yaml backend/app.yaml
'''

You can then check the front end on your [localhost:8080](http://localhost:8080)
And the generated Endpoints Apis on [http://localhost:8080/_ah/api/explorer](http://localhost:8080/_ah/api/explorer)
( To access to the endpoints, you may have to allow "Load Unsafe Script" in your browser.

## For more information

Google PaaS infrastructure [Go App Engine](https://cloud.google.com/appengine/docs/go/)
Google API Backend generator [Cloud Endpoints](https://cloud.google.com/endpoints/)
The frontend : [The Polymer Project](https://www.polymer-project.org)
Javascript package manager [Bower](http://bower.io/)
And our team [Business-Cloud](http://www.business-cloud.fr)


