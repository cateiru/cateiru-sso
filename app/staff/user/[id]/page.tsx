import React from 'react';
import {UserDetail} from '../../../../components/Staff/User/UserDetail';

interface Props {
  params: {
    id: string;
  };
}

const Page: React.FC<Props> = ({params}) => {
  return <UserDetail id={params.id} />;
};

export default Page;
