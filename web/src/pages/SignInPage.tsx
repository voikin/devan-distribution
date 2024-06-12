import React, { useState } from 'react';
import { AuthService } from '../services';
import { LoginRequest } from '../models/requests';
import { TextField, Button, Card, CardContent, CardActions, Typography, Box } from '@mui/material';
import { useAuthStore } from '../stores';
import { useMutation } from 'react-query';

export const SignInPage: React.FC = () => {
    const [credentials, setCredentials] = useState<LoginRequest>({
        username: '',
        password: '',
    });
    const login = useAuthStore((state) => state.login);

    const handleChange = (e: React.ChangeEvent<HTMLInputElement>) => {
        setCredentials({
            ...credentials,
            [e.target.name]: e.target.value,
        });
    };

    const loginMutation = useMutation(AuthService.login, {
        onSuccess: (data) => {
            login(data)
        }
    })

    const handleSubmit = async (e: React.FormEvent<HTMLFormElement>) => {
        e.preventDefault();
        loginMutation.mutate(credentials);
    };

    return (
        <Box pt={20} display="flex" justifyContent="center" alignItems="center">
            <Card>
                <CardContent>
                    <Typography variant="h5" component="div" gutterBottom>
                        Sign In
                    </Typography>
                    <form onSubmit={handleSubmit}>
                        <TextField
                            label="Username"
                            name="username"
                            value={credentials.username}
                            onChange={handleChange}
                            required
                            fullWidth
                            margin="normal"
                        />
                        <TextField
                            label="Password"
                            name="password"
                            type="password"
                            value={credentials.password}
                            onChange={handleChange}
                            required
                            fullWidth
                            margin="normal"
                        />
                        <CardActions>
                            <Button type="submit" variant="contained" color="primary" fullWidth>
                                Sign In
                            </Button>
                        </CardActions>
                    </form>
                </CardContent>
            </Card>
        </Box>
    );
};

export default SignInPage;
