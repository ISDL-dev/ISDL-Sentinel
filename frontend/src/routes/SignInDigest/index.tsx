import React, { useState } from 'react';
import axios from 'axios';
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
        title: 'Input required',
        description: 'Please enter both username and password',
        status: 'warning',
        duration: 5000,
        isClosable: true,
      });
      return;
    }

    try {
      const response = await axios.get(`${baseURL}/digest-login`, {
        auth: {
          username,
          password,
        },
        withCredentials: true,
      });

      const userData = response.data;
      setAuthUser(userData);

      toast({
        title: 'Login successful',
        description: `Welcome back, ${username}!`,
        status: 'success',
        duration: 5000,
        isClosable: true,
      });

      navigate('/');
    } catch (error) {
      console.error('Login error:', error);
      toast({
        title: 'Login failed',
        description: 'Invalid username or password',
        status: 'error',
        duration: 5000,
        isClosable: true,
      });
    }
  };

  const handleRegister = async () => {
    // ここに登録処理を実装します。
    // 実際の実装では、バックエンドAPIを呼び出してユーザーを登録します。
    toast({
      title: 'Register functionality',
      description: 'Register functionality is not implemented yet.',
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
            <Heading size={{ base: '1xl', md: "2xl" }}>Log in to your account</Heading>
            <Text color="gray.600">
              Don't have an account? <ChakraLink href="#">Sign up</ChakraLink>
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
                <FormLabel htmlFor="username">Username</FormLabel>
                <Input
                  id="username"
                  type="text"
                  value={username}
                  onChange={(e) => setUsername(e.target.value)}
                />
              </FormControl>
              <FormControl>
                <FormLabel htmlFor="password">Password</FormLabel>
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
                Register
              </Button>
              <Button colorScheme="teal" variant="solid" size="md" onClick={handleLogin}>
                Login
              </Button>
            </HStack>
          </Stack>
        </Box>
      </Stack>
    </Container>
  );
}