import UserDetails from '../../components/admin/UserDetails';
import Require from '../../components/common/Require';
import Title from '../../components/common/Title';

const AdminUser = () => {
  return (
    <Require isLogin={true} path="/" role="admin">
      <Title title="Admin | CateiruSSO" />
      <UserDetails />
    </Require>
  );
};

export default AdminUser;
