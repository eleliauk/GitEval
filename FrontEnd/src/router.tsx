import { createBrowserRouter } from 'react-router-dom';
import Login from './pages/login';
import Home from './pages/Home';
import { Evaluation } from './pages/Evaluation';

export const router = createBrowserRouter([
  {
    path: '/',
    element: <Login />,
  },
  {
    path: '/Home',
    element: <Home />,
  },
  {
    path: '/Eval',
    element: <Evaluation />,
  },
]);
