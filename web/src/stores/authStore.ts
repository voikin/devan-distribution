import { create } from 'zustand'
import { immer } from 'zustand/middleware/immer'
import { LoginResponse } from '../models/responses'
import { Role } from '../models'

interface UserState {
	role: Role | null
	isAuth: boolean
	login: (data: LoginResponse) => void
	logout: () => void
}

export const useAuthStore = create<UserState>()(
	immer((set) => ({
		role: null,
		isAuth: false,

		login: async (data: LoginResponse) => {
			localStorage.setItem('accessToken', data.accessToken)
			set((state) => {
				state.role = data.role
				state.isAuth = true
			})
		},

		logout: async () => {
			localStorage.removeItem('accessToken')
			set((state) => {
				state.role = null
				state.isAuth = false
			})
		},
	}))
)