import {
  ListItem,
  UnorderedList,
  useToast,
  Center,
  Box,
  Heading,
  Input,
  Button,
  Flex,
} from '@chakra-ui/react';
import NextLink from 'next/link';
import React from 'react';
import {IoArrowBackOutline} from 'react-icons/io5';
import {useSetRecoilState} from 'recoil';
import {
  getIPBanList,
  getMailBanList,
  setBan,
  deleteIPBan,
  deleteMailBan,
} from '../../utils/api/admin';
import {LoadState} from '../../utils/state/atom';

const BanPage = () => {
  const [ips, setIPs] = React.useState<string[]>([]);
  const [mails, setMails] = React.useState<string[]>([]);

  const [editIp, setEditIP] = React.useState('');
  const [editMail, setEditMail] = React.useState('');

  const toast = useToast();
  const setLoad = useSetRecoilState(LoadState);

  React.useEffect(() => {
    const f = async () => {
      try {
        setLoad(true);
        const ips = await getIPBanList();
        setIPs(ips);
        const mails = await getMailBanList();
        setMails(mails);
      } catch (error) {
        if (error instanceof Error) {
          toast({
            title: error.message,
            status: 'error',
            isClosable: true,
            duration: 9000,
          });
        }
      }

      setLoad(false);
    };

    f();
  }, []);

  const ipElement = (ip: string, i: number) => {
    return (
      <ListItem key={i} ml="1rem">
        {ip}
      </ListItem>
    );
  };

  const mailElement = (mail: string, i: number) => {
    return (
      <ListItem key={i} ml="1rem">
        {mail}
      </ListItem>
    );
  };

  const add = (mode: string) => {
    const f = async () => {
      try {
        switch (mode) {
          case 'ip':
            if (editIp === '') {
              return;
            }
            await setBan(editIp, undefined);
            setIPs(v => [...v, editIp]);
            setEditIP('');
            break;
          case 'mail':
            if (editMail === '') {
              return;
            }
            await setBan(undefined, editMail);
            setMails(v => [...v, editMail]);
            setEditMail('');
            break;
          default:
            break;
        }
      } catch (error) {
        if (error instanceof Error) {
          toast({
            title: error.message,
            status: 'error',
            isClosable: true,
            duration: 9000,
          });
        }
      }
    };

    f();
  };

  const deleteBan = (mode: string) => {
    const f = async () => {
      try {
        switch (mode) {
          case 'ip':
            if (editIp === '') {
              return;
            }
            await deleteIPBan(editIp);
            setIPs(v => {
              const index = v.indexOf(editIp);
              const e = [...v];
              e.splice(index, 1);
              return e;
            });
            setEditIP('');
            break;
          case 'mail':
            if (editMail === '') {
              return;
            }
            await deleteMailBan(editMail);
            setMails(v => {
              const index = v.indexOf(editMail);
              const e = [...v];
              e.splice(index, 1);
              return e;
            });
            setEditMail('');
            break;
          default:
            break;
        }
      } catch (error) {
        if (error instanceof Error) {
          toast({
            title: error.message,
            status: 'error',
            isClosable: true,
            duration: 9000,
          });
        }
      }
    };

    f();
  };

  return (
    <Center>
      <Box width={{base: '100%', lg: '1000px'}} mt="2rem">
        <Box mx=".5rem">
          <NextLink href="/setting/account" passHref>
            <Button
              pl=".5rem"
              variant="ghost"
              leftIcon={<IoArrowBackOutline size="25px" />}
            >
              戻る
            </Button>
          </NextLink>
        </Box>
        <Heading textAlign="center">Ban操作</Heading>
        <Box mx=".5rem" overflowX={{base: 'auto', lg: 'visible'}} mt="1rem">
          <Heading fontSize="1.5rem" mt="1rem">
            IPs
          </Heading>
          <UnorderedList>{ips.map(ipElement)}</UnorderedList>
          <Flex mt="1rem">
            <Input
              width="200px"
              mr=".5rem"
              value={editIp}
              onChange={v => setEditIP(v.target.value)}
            />
            <Button onClick={() => add('ip')}>追加</Button>
            <Button
              colorScheme="red"
              variant="ghost"
              onClick={() => deleteBan('ip')}
            >
              削除
            </Button>
          </Flex>

          <Heading fontSize="1.5rem" mt="1rem">
            Mails
          </Heading>
          <UnorderedList>{mails.map(mailElement)}</UnorderedList>
          <Flex mt="1rem">
            <Input
              width="200px"
              mr=".5rem"
              value={editMail}
              onChange={v => setEditMail(v.target.value)}
            />
            <Button onClick={() => add('mail')}>追加</Button>
            <Button
              colorScheme="red"
              variant="ghost"
              onClick={() => deleteBan('mail')}
            >
              削除
            </Button>
          </Flex>
        </Box>
      </Box>
    </Center>
  );
};

export default BanPage;
