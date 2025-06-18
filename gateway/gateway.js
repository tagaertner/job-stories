const { ApolloGateway, IntrospectAndCompose } = require('@apollo/gateway');
const { ApolloServer } = require('apollo-server-express');
const express = require('express');

async function startGateway() {
  console.log('🔄 Starting Apollo Federation Gateway...');

  const gateway = new ApolloGateway({
    supergraphSdl: new IntrospectAndCompose({
      subgraphs: [
        { name: 'products', url: 'http://localhost:4001/query' },
      ],
      pollIntervalInMs: 30000,
    }),
    debug: true,
  });

  const server = new ApolloServer({
    gateway,
    subscriptions: false,
    introspection: true,
    playground: true,
  });

  const app = express();
  
  await server.start();
  server.applyMiddleware({ app, path: '/graphql' });

  app.get('/health', (req, res) => {
    res.json({ 
      status: 'healthy', 
      gateway: 'apollo-federation',
    });
  });

  const port = 4000;
  
  app.listen(port, () => {
    console.log('🚀 Gateway ready at http://localhost:4000/graphql');
    console.log('🏥 Health check at http://localhost:4000/health');
  });
}

startGateway().catch(console.error);
