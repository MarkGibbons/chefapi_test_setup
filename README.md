# Chefapi Demo Setup
=================

Test setup uses the Chef Infra Server API to populate a Chef server with
organizations, users and nodes.  The users are added to the organizations
and to groups.  The nodes are assigned name and a dummy run list. 

This code is a fairly good example of how to use the Chef server API
via scripts. 

## Structure
The go modules have the code to add the objects.
The /bin scripts gather the credentials needed to talk to the Chef Server and run the go modules.
