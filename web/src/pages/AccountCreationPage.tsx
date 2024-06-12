import React, { useState } from 'react';
import { AuthService } from '../services';
import { CreateAccountRequest } from '../models/requests';
import { TextField, Button, MenuItem, Card, CardContent, CardActions, Typography, Box } from '@mui/material';
import { useMutation } from 'react-query';
import { useSnackbar } from 'notistack';
import { Role } from '../models';

export const AccountCreationPage: React.FC = () => {
    const [userData, setUserData] = useState<CreateAccountRequest>({
        username: '',
        password: '',
        role: Role.operator
    });

    const { enqueueSnackbar } = useSnackbar();

    const createAccountMutation = useMutation(AuthService.createAccount, {
        onSuccess: ({ id }) => {
            enqueueSnackbar(`Account created successfully with ID: ${id}`, { variant: 'success' });
        },
        onError: (error: Error) => {
            enqueueSnackbar(`Failed to create account: ${error.message}`, { variant: 'error' });
        }
    });

    const handleChange = (e: React.ChangeEvent<HTMLInputElement>) => {
        setUserData({
            ...userData,
            [e.target.name]: e.target.value as Role,
        });
    };

    const handleSubmit = async (e: React.FormEvent<HTMLFormElement>) => {
        e.preventDefault();
        createAccountMutation.mutate(userData);
    };

    return (
        <Box mt={20} display="flex" justifyContent="center" alignItems="center">
            <Card>
                <CardContent>
                    <Typography variant="h5" component="div" gutterBottom>
                        Create Account
                    </Typography>
                    <form onSubmit={handleSubmit}>
                        <TextField
                            label="Username"
                            name="username"
                            value={userData.username}
                            onChange={handleChange}
                            required
                            fullWidth
                            margin="normal"
                        />
                        <TextField
                            label="Password"
                            name="password"
                            type="password"
                            value={userData.password}
                            onChange={handleChange}
                            required
                            fullWidth
                            margin="normal"
                        />
                        <TextField
                            label="Role"
                            name="role"
                            select
                            value={userData.role}
                            onChange={handleChange}
                            required
                            fullWidth
                            margin="normal"
                        >
                            {Object.values(Role).map((role) => (
                                <MenuItem key={role} value={role}>
                                    {role}
                                </MenuItem>
                            ))}
                        </TextField>
                        <CardActions>
                            <Button type="submit" variant="contained" color="primary" fullWidth>
                                Create Account
                            </Button>
                        </CardActions>
                    </form>
                </CardContent>
            </Card>
        </Box>
    );
};

export default AccountCreationPage;
