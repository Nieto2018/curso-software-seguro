package models

// Necesario para time.Time

// PushEvent representa la estructura completa del JSON del webhook de GitHub
type PushEvent struct {
	Repository Repository `json:"repository"`
	HeadCommit *Commit    `json:"head_commit"` // Usamos *Commit porque puede ser null o no existir en ciertos eventos
}

// Repository representa la información del repositorio
type Repository struct {
	FullName string `json:"full_name"`
}

// Commit representa la información de un commit individual
type Commit struct {
	ID      string       `json:"id"`
	Message string       `json:"message"`
	Author  CommitAuthor `json:"author"`
}

// CommitAuthor representa la información del autor/comitter de un commit
type CommitAuthor struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

// Ejemplo de cómo usar las estructuras para hacer unmarshal
// func main() {
// 	jsonData := `
// 	{
// 		"ref": "refs/heads/main",
// 		"before": "17a2aba7c49616455fbec30b43dc09e74fb501e0",
// 		"after": "313070cd83d85c3b02e36435b122e0b14cd361ec",
// 		"repository": {
// 			"id": 999521414,
// 			"node_id": "R_kgDOO5N8hg",
// 			"name": "curso-software-seguro",
// 			"full_name": "Nieto2018/curso-software-seguro",
// 			"private": false,
// 			"owner": {
// 				"name": "Nieto2018",
// 				"email": "36169946+Nieto2018@users.noreply.github.com",
// 				"login": "Nieto2018",
// 				"id": 36169946,
// 				"node_id": "MDQ6VXNlcjM2MTY5OTQ2",
// 				"avatar_url": "https://avatars.github.com/u/36169946?v=4",
// 				"gravatar_id": "",
// 				"url": "https://api.github.com/users/Nieto2018",
// 				"html_url": "https://github.com/Nieto2018",
// 				"followers_url": "https://api.github.com/users/Nieto2018/followers",
// 				"following_url": "https://api.github.com/users/Nieto2018/following{/other_user}",
// 				"gists_url": "https://api.github.com/users/Nieto2018/gists{/gist_id}",
// 				"starred_url": "https://api.github.com/users/Nieto2018/starred{/owner}{/repo}",
// 				"subscriptions_url": "https://api.github.com/users/Nieto2018/subscriptions",
// 				"organizations_url": "https://api.github.com/users/Nieto2018/orgs",
// 				"repos_url": "https://api.github.com/users/Nieto2018/repos",
// 				"events_url": "https://api.github.com/users/Nieto2018/events{/privacy}",
// 				"received_events_url": "https://api.github.com/users/Nieto2018/received_events",
// 				"type": "User",
// 				"user_view_type": "public",
// 				"site_admin": false
// 			},
// 			"html_url": "https://github.com/Nieto2018/curso-software-seguro",
// 			"description": "https://platzi.com/cursos/software-seguro/",
// 			"fork": false,
// 			"url": "https://api.github.com/repos/Nieto2018/curso-software-seguro",
// 			"forks_url": "https://api.github.com/repos/Nieto2018/curso-software-seguro/forks",
// 			"keys_url": "https://api.github.com/repos/Nieto2018/curso-software-seguro/keys{/key_id}",
// 			"collaborators_url": "https://api.github.com/repos/Nieto2018/curso-software-seguro/collaborators{/collaborator}",
// 			"teams_url": "https://api.github.com/repos/Nieto2018/curso-software-seguro/teams",
// 			"hooks_url": "https://api.github.com/repos/Nieto2018/curso-software-seguro/hooks",
// 			"issue_events_url": "https://api.github.com/repos/Nieto2018/curso-software-seguro/issues/events{/number}",
// 			"events_url": "https://api.github.com/repos/Nieto2018/curso-software-seguro/events",
// 			"assignees_url": "https://api.github.com/repos/Nieto2018/curso-software-seguro/assignees{/user}",
// 			"branches_url": "https://api.github.com/repos/Nieto2018/curso-software-seguro/branches{/branch}",
// 			"tags_url": "https://api.github.com/repos/Nieto2018/curso-software-seguro/tags",
// 			"blobs_url": "https://api.github.com/repos/Nieto2018/curso-software-seguro/git/blobs{/sha}",
// 			"git_tags_url": "https://api.github.com/repos/Nieto2018/curso-software-seguro/git/tags{/sha}",
// 			"git_refs_url": "https://api.github.com/repos/Nieto2018/curso-software-seguro/git/refs{/sha}",
// 			"trees_url": "https://api.github.com/repos/Nieto2018/curso-software-seguro/git/trees{/sha}",
// 			"statuses_url": "https://api.github.com/repos/Nieto2018/curso-software-seguro/statuses/{sha}",
// 			"languages_url": "https://api.github.com/repos/Nieto2018/curso-software-seguro/languages",
// 			"stargazers_url": "https://api.github.com/repos/Nieto2018/curso-software-seguro/stargazers",
// 			"contributors_url": "https://api.github.com/repos/Nieto2018/curso-software-seguro/contributors",
// 			"subscribers_url": "https://api.github.com/repos/Nieto2018/curso-software-seguro/subscribers",
// 			"subscription_url": "https://api.github.com/repos/Nieto2018/curso-software-seguro/subscription",
// 			"commits_url": "https://api.github.com/repos/Nieto2018/curso-software-seguro/commits{/sha}",
// 			"git_commits_url": "https://api.github.com/repos/Nieto2018/curso-software-seguro/git/commits{/sha}",
// 			"comments_url": "https://api.github.com/repos/Nieto2018/curso-software-seguro/comments{/number}",
// 			"issue_comment_url": "https://api.github.com/repos/Nieto2018/curso-software-seguro/issues/comments{/number}",
// 			"contents_url": "https://api.github.com/repos/Nieto2018/curso-software-seguro/contents/{+path}",
// 			"compare_url": "https://api.github.com/repos/Nieto2018/curso-software-seguro/compare/{base}...{head}",
// 			"merges_url": "https://api.github.com/repos/Nieto2018/curso-software-seguro/merges",
// 			"archive_url": "https://api.github.com/repos/Nieto2018/curso-software-seguro/{archive_format}{/ref}",
// 			"downloads_url": "https://api.github.com/repos/Nieto2018/curso-software-seguro/downloads",
// 			"issues_url": "https://api.github.com/repos/Nieto2018/curso-software-seguro/issues{/number}",
// 			"pulls_url": "https://api.github.com/repos/Nieto2018/curso-software-seguro/pulls{/number}",
// 			"milestones_url": "https://api.github.com/repos/Nieto2018/curso-software-seguro/milestones{/number}",
// 			"notifications_url": "https://api.github.com/repos/Nieto2018/curso-software-seguro/notifications{?since,all,participating}",
// 			"labels_url": "https://api.github.com/repos/Nieto2018/curso-software-seguro/labels{/name}",
// 			"releases_url": "https://api.github.com/repos/Nieto2018/curso-software-seguro/releases{/id}",
// 			"deployments_url": "https://api.github.com/repos/Nieto2018/curso-software-seguro/deployments",
// 			"created_at": 1749555012,
// 			"updated_at": "2025-06-10T11:33:04Z",
// 			"pushed_at": 1749555440,
// 			"git_url": "git://github.com/Nieto2018/curso-software-seguro.git",
// 			"ssh_url": "git@github.com:Nieto2018/curso-software-seguro.git",
// 			"clone_url": "https://github.com/Nieto2018/curso-software-seguro.git",
// 			"svn_url": "https://github.com/Nieto2018/curso-software-seguro",
// 			"homepage": null,
// 			"size": 0,
// 			"stargazers_count": 0,
// 			"watchers_count": 0,
// 			"language": "Go",
// 			"has_issues": true,
// 			"has_projects": true,
// 			"has_downloads": true,
// 			"has_wiki": true,
// 			"has_pages": false,
// 			"has_discussions": false,
// 			"forks_count": 0,
// 			"mirror_url": null,
// 			"archived": false,
// 			"disabled": false,
// 			"open_issues_count": 0,
// 			"license": null,
// 			"allow_forking": true,
// 			"is_template": false,
// 			"web_commit_signoff_required": false,
// 			"topics": [],
// 			"visibility": "public",
// 			"forks": 0,
// 			"open_issues": 0,
// 			"watchers": 0,
// 			"default_branch": "main",
// 			"stargazers": 0,
// 			"master_branch": "main"
// 		},
// 		"pusher": {
// 			"name": "Nieto2018",
// 			"email": "36169946+Nieto2018@users.noreply.github.com"
// 		},
// 		"sender": {
// 			"login": "Nieto2018",
// 			"id": 36169946,
// 			"node_id": "MDQ6VXNlcjM2MTY5OTQ2",
// 			"avatar_url": "https://avatars.github.com/u/36169946?v=4",
// 			"gravatar_id": "",
// 			"url": "https://api.github.com/users/Nieto2018",
// 			"html_url": "https://github.com/Nieto2018",
// 			"followers_url": "https://api.github.com/users/Nieto2018/followers",
// 			"following_url": "https://api.github.com/users/Nieto2018/following{/other_user}",
// 			"gists_url": "https://api.github.com/users/Nieto2018/gists{/gist_id}",
// 			"starred_url": "https://api.github.com/users/Nieto2018/starred{/owner}{/repo}",
// 			"subscriptions_url": "https://api.github.com/users/Nieto2018/subscriptions",
// 			"organizations_url": "https://api.github.com/users/Nieto2018/orgs",
// 			"repos_url": "https://api.github.com/users/Nieto2018/repos",
// 			"events_url": "https://api.github.com/users/Nieto2018/events{/privacy}",
// 			"received_events_url": "https://api.github.com/users/Nieto2018/received_events",
// 			"type": "User",
// 			"user_view_type": "public",
// 			"site_admin": false
// 		},
// 		"created": false,
// 		"deleted": false,
// 		"forced": false,
// 		"base_ref": null,
// 		"compare": "https://github.com/Nieto2018/curso-software-seguro/compare/17a2aba7c496...313070cd83d8",
// 		"commits": [
// 			{
// 				"id": "313070cd83d85c3b02e36435b122e0b14cd361ec",
// 				"tree_id": "29bdbf123ad3593c2b9a38b47be11fdfe58fac8e",
// 				"distinct": true,
// 				"message": "test file",
// 				"timestamp": "2025-06-10T13:37:19+02:00",
// 				"url": "https://github.com/Nieto2018/curso-software-seguro/commit/313070cd83d85c3b02e36435b122e0b14cd361ec",
// 				"author": {
// 					"name": "Antonio Nieto García",
// 					"email": "antonio.nieto.garcia@externos.gloval.es"
// 				},
// 				"committer": {
// 					"name": "Antonio Nieto García",
// 					"email": "antonio.nieto.garcia@externos.gloval.es"
// 				},
// 				"added": [
// 					"test.txt"
// 				],
// 				"removed": [],
// 				"modified": []
// 			}
// 		],
// 		"head_commit": {
// 			"id": "313070cd83d85c3b02e36435b122e0b14cd361ec",
// 			"tree_id": "29bdbf123ad3593c2b9a38b47be11fdfe58fac8e",
// 			"distinct": true,
// 			"message": "test file",
// 			"timestamp": "2025-06-10T13:37:19+02:00",
// 			"url": "https://github.com/Nieto2018/curso-software-seguro/commit/313070cd83d85c3b02e36435b122e0b14cd361ec",
// 			"author": {
// 				"name": "Antonio Nieto García",
// 				"email": "antonio.nieto.garcia@externos.gloval.es"
// 			},
// 			"committer": {
// 				"name": "Antonio Nieto García",
// 				"email": "antonio.nieto.garcia@externos.gloval.es"
// 			},
// 			"added": [
// 				"test.txt"
// 			],
// 			"removed": [],
// 			"modified": []
// 		}
// 	}
// 	`

// 	var event PushEvent
// 	err := json.Unmarshal([]byte(jsonData), &event)
// 	if err != nil {
// 		fmt.Printf("Error al hacer unmarshal: %v\n", err)
// 		return
// 	}

// 	fmt.Printf("Ref: %s\n", event.Ref)
// 	fmt.Printf("Repository Name: %s\n", event.Repository.Name)
// 	fmt.Printf("Pusher Name: %s\n", event.Pusher.Name)
// 	if event.HeadCommit != nil {
// 		fmt.Printf("Head Commit Message: %s\n", event.HeadCommit.Message)
// 		fmt.Printf("Head Commit Author: %s\n", event.HeadCommit.Author.Name)
// 	} else {
// 		fmt.Println("No head commit found.")
// 	}

// 	fmt.Printf("First commit added files: %v\n", event.Commits[0].Added)
// }
