import AdminPage from '../components/admin/AdminPage';
import Require from '../components/common/Require';
import Title from '../components/common/Title';

const Admin = () => {
  return (
    <Require isLogin={true} path="/" role="admin">
      <Title title="Admin | CateiruSSO" />
      <AdminPage />
    </Require>
  );
};

export default Admin;
