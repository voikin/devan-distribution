import { Role } from "../role"

export interface LoginRequest {
	username: string
	password: string
}

export interface CreateAccountRequest {
	username: string
	password: string
	role: Role
}
