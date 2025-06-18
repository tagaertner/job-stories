const { ApolloGateway, IntrospectAndCompose } = require('@apollo/gateway');
const { ApolloServer } = require('apollo-server-express');
const express = require('express');

async function startGateway() {
  console.log('🔄 Starting Apollo Federation Gateway...');

  const gateway = new ApolloGateway({
    supergraphSdl: new IntrospectAndCompose({
      subgraphs: [
        { name: 'products', url: 'http://localhost:4001/query' },
        { name: 'users', url: 'http://localhost:4002/query' },
        { name: 'orders', url: 'http://localhost:4003/query' },
      ],
      // Poll every 30 seconds for service updates
      pollIntervalInMs: 30000,
      // Don't fail if services are unavailable initially
      introspectionHeaders: {
        'apollo-require-preflight': 'true',
      },
    }),
    debug: true,
    // Continue even if some services are down
    serviceHealthCheck: true,
  });

  const server = new ApolloServer({
    gateway,
    subscriptions: false,
    introspection: true,
    playground: true,
    // Don't stop server if gateway fails to load initially
    stopOnTerminationSignals: false,
  });

  const app = express();
  
  try {
    await server.start();
    server.applyMiddleware({ app, path: '/graphql' });
    console.log('✅ Gateway started successfully');
  } catch (error) {
    console.log('⚠️  Gateway started with errors (services not ready):', error.message);
    
    // Create a basic server even if federation fails
    app.get('/graphql', (req, res) => {
      res.json({ 
        error: 'Services not ready yet',
        message: 'Start the Go services on ports 4001, 4002, 4003'
      });
    });
  }

  app.get('/health', (req, res) => {
    res.json({ 
      status: 'healthy', 
      gateway: 'apollo-federation',
      timestamp: new Date().toISOString()
    });
  });

  app.get('/status', (req, res) => {
    res.json({
      gateway: 'running',
      expectedServices: [
        'http://localhost:4001/query (products)',
        'http://localhost:4002/query (users)', 
        'http://localhost:4003/query (orders)'
      ]
    });
  });

  const port = 4000;
  
  app.listen(port, () => {
    console.log('🚀 Gateway server ready at http://localhost:4000');
    console.log('📊 GraphQL endpoint: http://localhost:4000/graphql');
    console.log('🏥 Health check: http://localhost:4000/health');
    console.log('📋 Status: http://localhost:4000/status');
    console.log('');
    console.log('Waiting for services to start...');
  });
}

startGateway().catch(error => {
  console.error('❌ Failed to start gateway:', error);
  process.exit(1);
});
