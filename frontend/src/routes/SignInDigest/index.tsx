import React, { useState } from 'react';
import axios, { AxiosResponse } from 'axios';
import {
  Box,
  Button,
  Container,
  FormControl,
  FormLabel,
  Heading,
  HStack,
  Input,
  Stack,
  Text,
  useToast,
  Link as ChakraLink
} from '@chakra-ui/react';
import { useNavigate } from 'react-router-dom';
import { useUser } from '../../userContext';
import CryptoJS from 'crypto-js';

const baseURL = process.env.REACT_APP_BACKEND_ENDPOINT;

export default function SignInDigest() {
  const [username, setUsername] = useState('');
  const [password, setPassword] = useState('');
  const toast = useToast();
  const navigate = useNavigate();
  const { setAuthUser } = useUser();

  const handleLogin = async () => {
    if (username === '' || password === '') {
      toast({
        title: '入力が必要です',
        description: 'ユーザー名とパスワードの両方を入力してください',
        status: 'warning',
        duration: 5000,
        isClosable: true,
      });
      return;
    }

    try {
      // チャレンジを取得するための最初のリクエスト
      const challengeResponse: AxiosResponse = await axios.post(`${baseURL}/digest/login`, {}, {
        validateStatus: function (status) {
          return status < 500; // 500未満のステータスコードのみ解決
        }
      });

      if (challengeResponse.status === 401) {
        const wwwAuthenticateHeader: string | undefined = challengeResponse.headers['www-authenticate'];
        if (!wwwAuthenticateHeader) {
          throw new Error('WWW-Authenticate ヘッダーが見つかりません');
        }

        // WWW-Authenticate ヘッダーの解析
        const realm: string | undefined = wwwAuthenticateHeader.match(/realm="([^"]+)"/)?.[1];
        const nonce: string | undefined = wwwAuthenticateHeader.match(/nonce="([^"]+)"/)?.[1];
        const qop: string | undefined = wwwAuthenticateHeader.match(/qop="([^"]+)"/)?.[1];

        if (!realm || !nonce || !qop) {
          throw new Error('WWW-Authenticate ヘッダーの解析に失敗しました');
        }

        // レスポンスの生成
        const ha1: string = CryptoJS.MD5(`${username}:${realm}:${password}`).toString();
        const ha2: string = CryptoJS.MD5('POST:/digest/login').toString();
        const nc: string = '00000001';
        const cnonce: string = generateCnonce();
        const responseDigest: string = CryptoJS.MD5(`${ha1}:${nonce}:${nc}:${cnonce}:${qop}:${ha2}`).toString();

        // Authorization ヘッダーの構築
        const authHeader: string = `Digest username="${username}", realm="${realm}", nonce="${nonce}", uri="/digest/login", qop=${qop}, nc=${nc}, cnonce="${cnonce}", response="${responseDigest}", algorithm=MD5`;

        // Authorization ヘッダーを含む2回目のリクエスト
        const authResponse: AxiosResponse = await axios.post(`${baseURL}/digest/login`, {}, {
          headers: {
            'Authorization': authHeader
          }
        });

        const userData = authResponse.data;
        setAuthUser(userData);

        toast({
          title: 'ログイン成功',
          description: `おかえりなさい、${username}さん！`,
          status: 'success',
          duration: 5000,
          isClosable: true,
        });

        navigate('/');
      } else {
        throw new Error('サーバーからの予期しないレスポンス');
      }
    } catch (error) {
      console.error('ログインエラー:', error);
      toast({
        title: 'ログイン失敗',
        description: 'ユーザー名またはパスワードが無効です',
        status: 'error',
        duration: 5000,
        isClosable: true,
      });
    }
  };

  // クライアントノンス生成のためのヘルパー関数
  const generateCnonce = (): string => {
    return Math.random().toString(36).substring(2, 15) + Math.random().toString(36).substring(2, 15);
  };

  const handleRegister = async () => {
    // 登録機能の実装（必要な場合）
    toast({
      title: '登録機能',
      description: '登録機能はまだ実装されていません。',
      status: 'info',
      duration: 5000,
      isClosable: true,
    });
  };

  return (
    <Container maxW="lg" py={{ base: '32', md: '62' }} px={{ base: '0', sm: '8' }}>
      <Stack spacing="8">
        <Stack spacing="6">
          <Stack spacing={{ base: '2', md: '3' }} textAlign="center">
            <Heading size={{ base: '1xl', md: "2xl" }}>アカウントにログイン</Heading>
            <Text color="gray.600">
              アカウントをお持ちでない場合は <ChakraLink href="#">サインアップ</ChakraLink>
            </Text>
          </Stack>
        </Stack>
        <Box
          py={{ base: '8', sm: '8' }}
          px={{ base: '10', sm: '10' }}
          bg={{ base: 'white', sm: 'white' }}
          boxShadow={{ base: 'md', sm: 'md' }}
          borderRadius={{ base: 'xl', sm: 'xl' }}
        >
          <Stack spacing="6">
            <Stack spacing="5">
              <FormControl>
                <FormLabel htmlFor="username">ユーザー名</FormLabel>
                <Input
                  id="username"
                  type="text"
                  value={username}
                  onChange={(e) => setUsername(e.target.value)}
                />
              </FormControl>
              <FormControl>
                <FormLabel htmlFor="password">パスワード</FormLabel>
                <Input
                  id="password"
                  type="password"
                  value={password}
                  onChange={(e) => setPassword(e.target.value)}
                />
              </FormControl>
            </Stack>
            <HStack justify="space-between">
              <Button colorScheme="teal" variant="solid" size="md" onClick={handleRegister}>
                登録
              </Button>
              <Button colorScheme="teal" variant="solid" size="md" onClick={handleLogin}>
                ログイン
              </Button>
            </HStack>
          </Stack>
        </Box>
      </Stack>
    </Container>
  );
}