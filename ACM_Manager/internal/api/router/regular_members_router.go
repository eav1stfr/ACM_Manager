package router

import (
	"acmmanager/internal/api/handlers"
	"github.com/go-chi/chi/v5"
)

func registerRegularMembersRoutes(r chi.Router) {
	r.Route("/members", func(r chi.Router) {
		r.Get("/", handlers.GetMembersHandler) // get all members
		r.Post("/", handlers.CreateMembersHandler)
		r.Delete("/", handlers.DeleteMembersHandler)
		r.Patch("/", handlers.PatchMembersHandler)

		r.Route("/{id}", func(r chi.Router) {
			r.Get("/", handlers.GetOneMemberHandler)
			r.Patch("/", handlers.UpdateOneMemberHandler)
		})

		r.Route("/{department}", func(r chi.Router) {
			r.Get("/", handlers.GetMembersOfCertainDep)
		})
	})
}

//
//-- 1. Departments
//CREATE TABLE departments (
//name_of_dep TEXT PRIMARY KEY,
//head_id BIGINT UNIQUE
//);
//
//-- 2. DepHeads
//CREATE TABLE dep_heads (
//head_id BIGINT PRIMARY KEY,
//first_name TEXT NOT NULL,
//last_name TEXT NOT NULL,
//email TEXT UNIQUE NOT NULL,
//telegram TEXT,
//role TEXT,
//dep_id TEXT REFERENCES departments(name_of_dep) ON DELETE SET NULL
//);
//
//-- 3. Board Members
//CREATE TABLE board_members (
//id BIGINT PRIMARY KEY,
//first_name TEXT NOT NULL,
//last_name TEXT NOT NULL,
//email TEXT UNIQUE NOT NULL,
//telegram TEXT,
//role TEXT
//);
//
//-- 4. Regular Members
//CREATE TABLE regular_members (
//id BIGINT PRIMARY KEY,
//first_name TEXT NOT NULL,
//last_name TEXT NOT NULL,
//email TEXT UNIQUE NOT NULL,
//telegram TEXT,
//role TEXT
//);
//
//-- 5. Department-Member Join Table
//CREATE TABLE department_members (
//member_id BIGINT REFERENCES regular_members(id) ON DELETE CASCADE,
//department_name TEXT REFERENCES departments(name_of_dep) ON DELETE CASCADE,
//PRIMARY KEY (member_id, department_name)
//);
//
//-- 6. Tasks
//CREATE TABLE tasks (
//id BIGINT PRIMARY KEY,
//description TEXT NOT NULL,
//deadline TIMESTAMP NOT NULL,
//complexity INT CHECK (complexity IN (0, 1, 2)), -- 0=Easy, 1=Medium, 2=Hard
//status BOOLEAN DEFAULT FALSE,
//assigned BOOLEAN DEFAULT FALSE,
//assigned_to BIGINT REFERENCES regular_members(id) ON DELETE SET NULL
//);
//
//-- 7. Member-Task Join Table (for tasks done)
//CREATE TABLE member_tasks_done (
//member_id BIGINT REFERENCES regular_members(id) ON DELETE CASCADE,
//task_id BIGINT REFERENCES tasks(id) ON DELETE CASCADE,
//PRIMARY KEY (member_id, task_id)
//);
//
//-- 8. Member-Task Join Table (for tasks to do)
//CREATE TABLE member_tasks_todo (
//member_id BIGINT REFERENCES regular_members(id) ON DELETE CASCADE,
//task_id BIGINT REFERENCES tasks(id) ON DELETE CASCADE,
//PRIMARY KEY (member_id, task_id)
//);
//
//-- 9. Meetings
//CREATE TABLE meetings (
//id BIGINT PRIMARY KEY,
//venue TEXT NOT NULL,
//time TIMESTAMP NOT NULL,
//day INT CHECK (day BETWEEN 0 AND 6), -- Sunday = 0
//repeated BOOLEAN DEFAULT FALSE,
//for_department BOOLEAN DEFAULT FALSE,
//department_name TEXT REFERENCES departments(name_of_dep) ON DELETE SET NULL
//);
