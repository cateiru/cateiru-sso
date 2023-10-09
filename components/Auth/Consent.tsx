import {
  Accordion,
  AccordionButton,
  AccordionIcon,
  AccordionItem,
  AccordionPanel,
  Box,
  Button,
  Center,
  Flex,
  Heading,
  Link,
  ListItem,
  Stack,
  Text,
  UnorderedList,
  keyframes,
} from '@chakra-ui/react';
import React from 'react';
import {
  TbExternalLink,
  TbFileDescription,
  TbPlugConnected,
  TbSquareKey,
  TbUserQuestion,
} from 'react-icons/tb';
import {validateScope} from '../../utils/validate';
import {Avatar} from '../Common/Chakra/Avatar';
import {useShadowColor} from '../Common/useColor';

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
  userImage?: string;

  clientName: string;
  description?: string;
  clientImage?: string;

  registerUserName: string;
  registerUserImage?: string;

  orgName?: string;
  orgImage?: string;
  orgMemberOnly: boolean;

  scopes?: string[];
  redirectUri: string;

  onSubmit: () => Promise<void>;
  onCancel: () => Promise<void>;
}

export const Consent: React.FC<Props> = props => {
  const shadowColor = useShadowColor();
  const [isLoadingSubmit, setIsLoadingSubmit] = React.useState(false);
  const [isLoadingCancel, setIsLoadingCancel] = React.useState(false);

  const redirectUriHost = new URL(props.redirectUri);
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
      boxShadow={{base: 'none', sm: `0px 0px 7px -2px ${shadowColor}`}}
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
          {props.clientName}
        </Heading>
      </Center>
      <Text textAlign="center" fontWeight="bold">
        がアクセスを求めています
      </Text>

      <Center mt="1.5rem">
        <Avatar src={props.userImage} justifyContent="flex-start" size="lg" />
        <Center
          mx="40px"
          cursor="pointer"
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
        <Avatar src={props.clientImage} justifyContent="flex-end" size="lg" />
      </Center>

      {props.orgMemberOnly && (
        <Text textAlign="center" my=".5rem" fontSize=".8rem">
          このクライアントは組織に所属しているユーザーのみ有効です。
        </Text>
      )}

      <Accordion mt="1rem" defaultIndex={[0]} allowMultiple>
        {props.description && (
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
                {props.description}
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
              {props.scopes?.map(v => {
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
                {`作成者${props.orgName ? '・組織' : ''}に関する情報`}
              </Text>
              <AccordionIcon />
            </AccordionButton>
          </Text>
          <AccordionPanel>
            {props.orgName && (
              <Text mb="1rem">
                ※このクライアントは、組織
                <Text as="span" fontWeight="bold" lineHeight="32px" mx=".5rem">
                  <Avatar src={props.orgImage} size="sm" mr=".2rem" />
                  {props.orgName}
                </Text>
                によって管理されています。
              </Text>
            )}
            <Flex alignItems="center">
              <Avatar src={props.registerUserImage} />
              <Text ml=".5rem" fontWeight="bold">
                {props.registerUserName}
              </Text>
            </Flex>
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
        <Text mt=".5rem">
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
