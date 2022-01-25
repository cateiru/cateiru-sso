import Title from '../components/common/Title';
import ShowUser from '../components/user/ShowUser';

const User = () => {
  return (
    <>
      <Title title="ユーザ | CateiruSSO" />
      <style>
        {`@import
        url('https://fonts.googleapis.com/css2?family=Source+Code+Pro:wght@500;600&display=swap');`}
      </style>
      <ShowUser />
    </>
  );
};

export default User;
