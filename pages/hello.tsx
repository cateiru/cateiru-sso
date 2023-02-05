import NoSSR from 'react-no-ssr';
import Title from '../components/common/Title';
import ShowUser from '../components/user/ShowUser';

const UserHello = () => {
  return (
    <>
      <Title title="ユーザ | CateiruSSO" />
      <NoSSR>
        <style>
          {
            "@import url('https://fonts.googleapis.com/css2?family=Source+Code+Pro:wght@500;600&display=swap');"
          }
        </style>
      </NoSSR>
      <ShowUser />
    </>
  );
};

export default UserHello;
