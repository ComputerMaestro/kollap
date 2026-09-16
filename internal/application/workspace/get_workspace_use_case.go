package workspace

// func GetWorkspace(w http.ResponseWriter, r *http.Request) {
// 	ctx := r.Context()

// 	var p GetWorkspaceRequest
// 	decoder := json.NewDecoder(r.Body)
// 	err := decoder.Decode(&p)
// 	if err != nil {
// 		log.Printf("invalid payload - %s", err)
// 	}

// 	workspace, err := WorkspaceDao.FindByID(ctx, p.ID)
// 	if err != nil {
// 		log.Print(fmt.Errorf("error finding workspace %v", err))
// 	}

// 	w.WriteHeader(http.StatusFound)
// 	bytes, _ := json.Marshal(&Workspace{
// 		ID:        workspace.ID,
// 		Name:      workspace.Name,
// 		CreatedAt: workspace.CreatedAt,
// 		UpdatedAt: workspace.UpdatedAt,
// 	})
// 	w.Write(bytes)
// }
