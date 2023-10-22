import {
  Accordion,
  AccordionButton,
  AccordionIcon,
  AccordionItem,
  AccordionPanel,
  Badge,
  Box,
  Button,
  Center,
  Flex,
  Heading,
  Link,
  ListItem,
  Stack,
  Table,
  TableContainer,
  Tbody,
  Td,
  Text,
  Tr,
  UnorderedList,
  keyframes,
  useColorModeValue,
} from '@chakra-ui/react';
import React from 'react';
import {
  TbDiscountCheck,
  TbExternalLink,
  TbFileDescription,
  TbPlugConnected,
  TbSquareKey,
  TbUserQuestion,
} from 'react-icons/tb';
import {PublicAuthenticationRequest} from '../../utils/types/auth';
import {validateScope} from '../../utils/validate';
import {Avatar} from '../Common/Chakra/Avatar';
import {Tooltip} from '../Common/Chakra/Tooltip';

const rotate = keyframes`
0%,
100% {
  transform: rotate(0deg);
}
100% {
  transform: rotate(360deg);
}
`;

interface Props {
  userName: string;
  userImage?: string;

  data: PublicAuthenticationRequest;

  onSubmit: () => Promise<void>;
  onCancel: () => Promise<void>;
}

export const Consent: React.FC<Props> = props => {
  const shadow = useColorModeValue(
    '0px 0px 5px -2px #242838',
    '0px 0px 10px -2px #000000'
  );
  const [isLoadingSubmit, setIsLoadingSubmit] = React.useState(false);
  const [isLoadingCancel, setIsLoadingCancel] = React.useState(false);

  const redirectUriHost = new URL(props.data.redirect_uri);
  redirectUriHost.pathname = '';

  const onSubmit = () => {
    setIsLoadingSubmit(true);

    props
      .onSubmit()
      .then(() => {
        setIsLoadingSubmit(false);
      })
      .catch(() => {
        setIsLoadingSubmit(false);
      });
  };

  const onCancel = () => {
    setIsLoadingCancel(true);

    props
      .onCancel()
      .then(() => {
        setIsLoadingCancel(false);
      })
      .catch(() => {
        setIsLoadingCancel(false);
      });
  };

  return (
    <Flex
      w={{base: '96%', sm: '450px'}}
      minH={{base: '600px', sm: '700px'}}
      boxShadow={shadow}
      borderRadius="10px"
      mx="auto"
      py="2rem"
      flexDirection="column"
    >
      <Center>
        <Heading
          fontWeight="bold"
          textAlign="center"
          background="linear-gradient(132deg, #e23d3d 0%, #ec44bd 100%);"
          backgroundClip="text"
        >
          {props.data.client_name}
        </Heading>
      </Center>
      <Text textAlign="center" fontWeight="bold">
        がアクセスを求めています
      </Text>

      <Center mt="1.5rem">
        <Tooltip label={props.userName} placement="top">
          <Avatar src={props.userImage} justifyContent="flex-start" size="lg" />
        </Tooltip>
        <Center
          mx="40px"
          css={{
            ':hover': {
              animation: `${rotate} 1s linear infinite`,
            },
          }}
        >
          <TbPlugConnected
            size="40px"
            style={{
              transform: 'rotate(45deg)',
            }}
          />
        </Center>
        <Tooltip label={props.data.client_name} placement="top">
          <Avatar
            src={props.data.image ?? ''}
            justifyContent="flex-end"
            size="lg"
          />
        </Tooltip>
      </Center>

      {props.data.org_member_only && (
        <Text textAlign="center" my=".5rem" fontSize=".8rem">
          このクライアントは組織に所属しているユーザーのみ有効です。
        </Text>
      )}

      <Accordion
        mt="1rem"
        defaultIndex={props.data.client_description ? [0, 1] : [0]}
        allowMultiple
      >
        {props.data.client_description && (
          <AccordionItem>
            <Text as="h2">
              <AccordionButton>
                <TbFileDescription size="20px" />
                <Text
                  ml=".5rem"
                  as="span"
                  flex="1"
                  textAlign="left"
                  fontWeight="bold"
                >
                  クライアントの説明
                </Text>
                <AccordionIcon />
              </AccordionButton>
            </Text>
            <AccordionPanel mx=".5rem">
              <Text as="pre" whiteSpace="pre-wrap" w="100%">
                {props.data.client_description}
              </Text>
            </AccordionPanel>
          </AccordionItem>
        )}
        <AccordionItem>
          <Text as="h2">
            <AccordionButton>
              <TbSquareKey size="20px" />
              <Text
                ml=".5rem"
                as="span"
                flex="1"
                textAlign="left"
                fontWeight="bold"
              >
                求めている権限
              </Text>
              <AccordionIcon />
            </AccordionButton>
          </Text>
          <AccordionPanel>
            <UnorderedList pl="1rem" mb=".5rem">
              {props.data.scopes?.map(v => {
                return (
                  <ListItem key={`scope-${v}`}>{validateScope(v)}</ListItem>
                );
              })}
            </UnorderedList>
          </AccordionPanel>
        </AccordionItem>
        <AccordionItem>
          <Text as="h2">
            <AccordionButton>
              <TbUserQuestion size="20px" />
              <Text
                ml=".5rem"
                as="span"
                flex="1"
                textAlign="left"
                fontWeight="bold"
              >
                {`作成者${props.data.org_name ? '・組織' : ''}に関する情報`}
              </Text>
              <AccordionIcon />
            </AccordionButton>
          </Text>
          <AccordionPanel>
            {props.data.org_name && (
              <Text mb="1rem">
                ※このクライアントは、組織
                <Text as="span" fontWeight="bold" lineHeight="32px" mx=".5rem">
                  <Avatar
                    src={props.data.org_image ?? ''}
                    size="sm"
                    mr=".2rem"
                  />
                  {props.data.org_name}
                </Text>
                によって管理されています。
              </Text>
            )}
            <Flex alignItems="center">
              <Avatar src={props.data.register_user_image ?? ''} size="sm" />
              <Text ml=".5rem" fontWeight="bold">
                {props.data.register_user_name}
              </Text>
            </Flex>
          </AccordionPanel>
        </AccordionItem>
        <AccordionItem borderBottom="none">
          <Text as="h2">
            <AccordionButton>
              <TbDiscountCheck size="20px" />
              <Text
                ml=".5rem"
                as="span"
                flex="1"
                textAlign="left"
                fontWeight="bold"
              >
                OpenIDに関する情報
              </Text>
              <AccordionIcon />
            </AccordionButton>
          </Text>
          <AccordionPanel px="0" pb="0">
            <TableContainer>
              <Table variant="simple">
                <Tbody>
                  <Tr>
                    <Td fontWeight="bold">クライアントID</Td>
                    <Td>{props.data.client_id}</Td>
                  </Tr>
                  <Tr>
                    <Td fontWeight="bold">レスポンスタイプ</Td>
                    <Td>{props.data.response_type}</Td>
                  </Tr>
                  <Tr>
                    <Td fontWeight="bold">プロンプト</Td>
                    <Td>
                      {props.data.prompts.map(v => {
                        return (
                          <Badge
                            key={`prompts-${v}`}
                            colorScheme="cateiru"
                            mr=".5rem"
                          >
                            {v}
                          </Badge>
                        );
                      })}
                    </Td>
                  </Tr>
                </Tbody>
              </Table>
            </TableContainer>
          </AccordionPanel>
        </AccordionItem>
      </Accordion>

      <Box mt="auto" pt="1rem" w={{base: '100%', sm: '96%'}} mx="auto">
        <Stack direction="column" spacing="10px">
          <Button
            w="100%"
            colorScheme="cateiru"
            isLoading={isLoadingSubmit}
            isDisabled={isLoadingCancel}
            onClick={onSubmit}
          >
            許可する
          </Button>
          <Button
            isDisabled={isLoadingSubmit || isLoadingCancel}
            onClick={onCancel}
            variant="link"
          >
            キャンセル
          </Button>
        </Stack>
        <Text mt="1rem">
          許可すると、
          <Link
            href={redirectUriHost.toString()}
            isExternal
            mx=".25rem"
            fontWeight="bold"
          >
            {redirectUriHost.host}
            <TbExternalLink
              size="1rem"
              style={{
                display: 'inline-block',
                verticalAlign: 'middle',
                marginLeft: '0.1rem',
              }}
            />
          </Link>
          にリダイレクトします
        </Text>
      </Box>
    </Flex>
  );
};
