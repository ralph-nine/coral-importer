package common

import "gitlab.com/coralproject/coral-importer/common/coral"

type Reconstructor struct {
	parents  map[string]string
	children map[string][]string
}

func (r *Reconstructor) AddComment(comment *coral.Comment) {
	// Save the comment's parent.
	r.parents[comment.ID] = comment.ParentID

	// Add the comment to the array of children, if it has any.
	if comment.ParentID != "" {
		if _, ok := r.children[comment.ParentID]; !ok {
			r.children[comment.ParentID] = []string{comment.ID}
		} else {
			r.children[comment.ParentID] = append(r.children[comment.ParentID], comment.ID)
		}
	}
}

// GetChildren will get the linked children for a Comment.
func (r *Reconstructor) GetChildren(commentID string) []string {
	children, ok := r.children[commentID]
	if ok {
		return children
	}

	return []string{}
}

// GetParent will get a parent (if it exists) for a Comment.
func (r *Reconstructor) GetParent(commentID string) string {
	parentID, ok := r.parents[commentID]
	if ok && parentID != "" {
		return parentID
	}

	return ""
}

// GetAncestors will get the array of ancestors for a given Comment.
func (r *Reconstructor) GetAncestors(commentID string) []string {
	// Store the ancestors in an array.
	ancestorIDs := make([]string, 0)

	parentID := r.GetParent(commentID)
	for parentID != "" {
		ancestorParentID, ok := r.parents[parentID]
		if !ok {
			return nil
		}

		// Add the ID of this comment to the ancestorIDs.
		ancestorIDs = append(ancestorIDs, parentID)

		// Store the reference to the ancestorParentID
		parentID = ancestorParentID
	}

	return ancestorIDs
}

func NewReconstructor() *Reconstructor {
	return &Reconstructor{
		parents:  make(map[string]string),
		children: make(map[string][]string),
	}
}
