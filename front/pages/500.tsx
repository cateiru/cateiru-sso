import Title from '../components/common/Title';
import InternalServerPage from '../components/error/InternalServerPage';

const InternalServer = () => {
  return (
    <>
      <Title title="500 | CateiruSSO" />
      <InternalServerPage />
    </>
  );
};

export default InternalServer;
