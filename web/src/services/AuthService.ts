import { CreateAccountResponse, LoginResponse } from './../models/responses'
import axios from 'axios'
import { $api, API_URL } from '../http'
import { CreateAccountRequest, LoginRequest } from '../models/requests'

export const AuthService = {
	async login(credentials: LoginRequest): Promise<LoginResponse> {
		try {
			const response = await $api.post<LoginResponse>('/auth/login', credentials)
			return response.data
		} catch (e) {
			console.log(e)
            throw e
		}
	},

	async createAccount(credentials: CreateAccountRequest): Promise<CreateAccountResponse> {
		try {
			const response = await $api.post<CreateAccountResponse>('/auth/create-account', credentials)
			return response.data
		} catch (e) {
			console.log(e)
            throw e
		}
	},

	async logout(): Promise<void> {
		return $api.get('/auth/logout')
	},

	async checkAuth() {
        try {
            const response = await axios.get<LoginResponse>(
                `${API_URL}/auth/refresh`, 
                { withCredentials: true }
            )
            return response.data
		} catch (e) {
			console.log(e)
            throw e
		}
	}
}
