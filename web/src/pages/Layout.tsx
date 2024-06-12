import React from 'react';
import { AppBar, Toolbar, Typography, Button, Container, Box } from '@mui/material';
import { useAuthStore } from '../stores';
import { AuthService } from '../services';
import { useMutation } from 'react-query';

// eslint-disable-next-line @typescript-eslint/no-explicit-any
export const Layout: React.FC<{ children: any }> = ({ children }) => {
    const isAuth = useAuthStore((state) => state.isAuth);
    const logout = useAuthStore((state) => state.logout);

    const logoutMutation = useMutation(AuthService.logout, {
        onSuccess: () => {
            logout();
        }
    });

    return (
        <Box display="flex" flexDirection="column" minHeight="100vh">
            <AppBar position="static">
                <Toolbar>
                    <Typography variant="h6" style={{ flexGrow: 1 }}>
                        Devan Distribution
                    </Typography>
                    {isAuth && (
                        <Button color="inherit" onClick={() => logoutMutation.mutate()}>
                            Logout
                        </Button>
                    )}
                </Toolbar>
            </AppBar>
            <Container component="main" style={{ flexGrow: 1 }}>
                {children}
            </Container>
            <Box
                component="footer"
                sx={{
                    py: 2,
                    px: 2,
                    mt: 'auto',
                    backgroundColor: (theme) =>
                        theme.palette.mode === 'light' ? '#f1f1f1' : theme.palette.grey[800],
                    textAlign: 'center'
                }}
            >
                <Typography variant="body2" color="textSecondary">
                    &copy; 2023 Devan Distribution. All rights reserved.
                </Typography>
            </Box>
        </Box>
    );
};
