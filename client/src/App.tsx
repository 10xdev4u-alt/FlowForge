import { Routes, Route, Navigate } from 'react-router-dom';
import { useAuthStore } from './store/authStore';
import { Login } from './pages/Login';
import { Register } from './pages/Register';
import { Dashboard } from './pages/Dashboard';
import { CreateWorkflow } from './pages/CreateWorkflow';
import { Editor } from './pages/Editor';

function App() {
  const token = useAuthStore((state) => state.token);

  return (
    <Routes>
      <Route path="/login" element={<Login />} />
      <Route path="/register" element={<Register />} />
      <Route path="/" element={token ? <Dashboard /> : <Navigate to="/login" />} />
      <Route path="/workflows/new" element={token ? <CreateWorkflow /> : <Navigate to="/login" />} />
      <Route path="/editor/:id" element={token ? <Editor /> : <Navigate to="/login" />} />
    </Routes>
  );
}

export default App;
