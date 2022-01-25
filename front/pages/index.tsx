import type {NextPage} from 'next';
import Title from '../components/common/Title';
import TopPage from '../components/top/TopPage';

const Home: NextPage = () => {
  return (
    <>
      <Title title="Top | CateiruSSO" />
      <TopPage />
    </>
  );
};

export default Home;
