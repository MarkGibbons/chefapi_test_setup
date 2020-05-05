package testapi

import (
	"fmt"
	"os"

	"github.com/go-chef/chef"
)

// Create users
// Add the users to the organizations
func User() {
	orgs := []string{"blue", "sales", "admin", "testorg", "mainapp"}
	users := []string{"arline", "doan", "fred", "sue"}
	client := Client()
	for _, user := range users {
		usr := chef.User{
			UserName:    user,
			Email:       user + "@domain.io",
			FirstName:   user,
			LastName:    "fullname",
			DisplayName: user + "Fullname",
			Password:    "Logn12ComplexPwd#",
			CreateKey:   true,
		}
		newuser, err := client.Users.Create(usr)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Couldn't create user %+v Err: %+v\n", user, err)
		}
		fmt.Printf("Created %+v result %+v err %+v\n", usr.UserName, newuser, err)
	}

	for _, org := range orgs {
		client := OrgClient(org)
		for _, user := range users {
			adduser := chef.AddNow{Username: user}
			err := client.Associations.Add(adduser)
			if err != nil {
				fmt.Fprintf(os.Stderr, "Couldn't add user %+v to org %+v Err: %+v\n", user, org, err)
			}
		}
	}

	// sue is admin for these organizations
	sueadmins := []string{"blue", "sales", "admin", "mainapp"}
	for _, org := range sueadmins {
		client = OrgClient(org)
		admingrp, err := client.Groups.Get("admins")
		// Add admin users
		groupupdate := chef.GroupUpdate{}
		groupupdate.Name = "admins"
		groupupdate.GroupName = "admins"
		groupupdate.Actors.Clients = admingrp.Clients
		groupupdate.Actors.Groups = admingrp.Groups
		groupupdate.Actors.Users = append(admingrp.Users, "sue")
		_, err = client.Groups.Update(groupupdate)
		adminout, _ := client.Groups.Get("admins")
		fmt.Printf("For org %+v admins %+v err  %+v\n", org, adminout, err)
	}

	// doan is admin for the testorg organization
	client = OrgClient("testorg")
	// Get the existing group settings
	admingrp, err := client.Groups.Get("admins")
	// Create an update
	groupupdate := chef.GroupUpdate{}
	groupupdate.Name = "admins"
	groupupdate.GroupName = "admins"
	groupupdate.Actors.Clients = admingrp.Clients
	groupupdate.Actors.Groups = admingrp.Groups
	groupupdate.Actors.Users = append(admingrp.Users, "doan")
	_, err = client.Groups.Update(groupupdate)
	adminout, _ := client.Groups.Get("admins")
	fmt.Printf("For org %+v admins %+v err  %+v\n", "testorg", adminout, err)
}
