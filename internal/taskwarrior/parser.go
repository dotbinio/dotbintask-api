package taskwarrior

import (
	"strings"
)

// FilterTasks applies filters to a list of tasks
func FilterTasks(tasks []Task, filter TaskFilter) []Task {
	filtered := make([]Task, 0)

	for _, task := range tasks {
		if !matchesFilter(task, filter) {
			continue
		}
		filtered = append(filtered, task)
	}

	return filtered
}

// matchesFilter checks if a task matches the given filter
func matchesFilter(task Task, filter TaskFilter) bool {
	// Filter by status
	if filter.Status != "" && task.Status != filter.Status {
		return false
	}

	// Filter by project
	if filter.Project != "" && task.Project != filter.Project {
		return false
	}

	// Filter by UUID
	if filter.UUID != "" && task.UUID != filter.UUID {
		return false
	}

	// Filter by tags (task must have ALL specified tags)
	if len(filter.Tags) > 0 {
		taskTagSet := make(map[string]bool)
		for _, tag := range task.Tags {
			taskTagSet[tag] = true
		}

		for _, requiredTag := range filter.Tags {
			if !taskTagSet[requiredTag] {
				return false
			}
		}
	}

	return true
}

// ExtractProjectsFromTasks extracts unique projects from tasks
func ExtractProjectsFromTasks(tasks []Task) []string {
	projectSet := make(map[string]bool)
	for _, task := range tasks {
		if task.Project != "" {
			projectSet[task.Project] = true
		}
	}

	projects := make([]string, 0, len(projectSet))
	for project := range projectSet {
		projects = append(projects, project)
	}

	return projects
}

// ExtractTagsFromTasks extracts unique tags from tasks
func ExtractTagsFromTasks(tasks []Task) []string {
	tagSet := make(map[string]bool)
	for _, task := range tasks {
		for _, tag := range task.Tags {
			tagSet[tag] = true
		}
	}

	tags := make([]string, 0, len(tagSet))
	for tag := range tagSet {
		tags = append(tags, tag)
	}

	return tags
}

// ValidateTaskUUID checks if a UUID looks valid
func ValidateTaskUUID(uuid string) bool {
	// Basic validation - Taskwarrior UUIDs are standard UUIDs
	// Format: xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx
	if len(uuid) != 36 {
		return false
	}

	if uuid[8] != '-' || uuid[13] != '-' || uuid[18] != '-' || uuid[23] != '-' {
		return false
	}

	return true
}

// SanitizeInput sanitizes user input to prevent command injection
func SanitizeInput(input string) string {
	// Remove potentially dangerous characters
	replacer := strings.NewReplacer(
		";", "",
		"&", "",
		"|", "",
		"`", "",
		"$", "",
		"(", "",
		")", "",
		"<", "",
		">", "",
		"\n", "",
		"\r", "",
	)
	return replacer.Replace(input)
}
