import React, { useEffect } from 'react';
import { BrowserRouter as Router, Route, Routes } from 'react-router-dom';
import { useMutation } from 'react-query';
import { useAuthStore } from './stores';
import { AuthService } from './services/AuthService';
import { AccountCreationPage, HomePage, Layout, OrdersPage, SignInPage } from './pages';
import { Role } from './models';

export const App: React.FC = () => {
    const {isAuth, login, role } = useAuthStore();

    const loginMutation = useMutation(AuthService.checkAuth, {
        onSuccess: (data) => {
            login(data);
        },
        onError: (error) => {
            console.log(error);
        },
    });

    useEffect(() => {
        const accessToken = localStorage.getItem('accessToken');
        if (accessToken) loginMutation.mutate();
    // eslint-disable-next-line react-hooks/exhaustive-deps
    }, []);

    return (
        <Router>
            <Layout>
                <Routes>
                    {/* common pages */}
                    <Route path="/" element={<HomePage />} />
                    {!isAuth && <Route path="/sign-in" element={<SignInPage />} />}

                    {/* specialist pages */}
                    {isAuth && <Route path="/orders" element={<OrdersPage/>}/>}

                    {/* admin pages  */}
                    {!isAuth && role !== Role.admin && <Route path="/create-account" element={<AccountCreationPage />} />}
                </Routes>
            </Layout>
        </Router>
    );
};
