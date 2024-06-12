import { Role } from "../role"

export interface LoginResponse {
	role: Role
	accessToken: string
}

export interface CreateAccountResponse {
	id: number
}
