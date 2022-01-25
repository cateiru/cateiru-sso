import {Text, Flex} from '@chakra-ui/react';

const NotFoundPage = () => {
  return (
    <Flex
      justifyContent="center"
      alignItems="center"
      flexDirection="column"
      width="100%"
      height="80vh"
      px={{base: '1rem', md: '5rem'}}
    >
      <Text fontSize="2.5rem" fontWeight="light" marginBottom="1.5rem">
        404 | Not Found.
      </Text>
      <Text fontSize="1.5rem" fontWeight="light">
        ﾊﾞﾝﾊﾞﾝﾊﾞﾝﾊﾞﾝﾊﾞﾝﾊﾞﾝﾊﾞﾝ
        <br />
        ﾊﾞﾝ&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;ﾊﾞﾝﾊﾞﾝﾊﾞﾝ
        <br />
        ﾊﾞﾝ&nbsp;&nbsp;(∩`･ω･)&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;ﾊﾞﾝﾊﾞﾝ
        <br />
        &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;＿/_ﾐつ/￣￣￣/
        <br />
        &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;＼/＿＿＿/￣￣
        <br />
      </Text>
    </Flex>
  );
};

export default NotFoundPage;
