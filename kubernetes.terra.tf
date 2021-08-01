resource "kubernetes_config_map" "pic_it_configmap" {
  metadata {
    name = "pic-it-configmap"
  }

  data = {
    HOST = "mysql-service"

    PORT = "3306"

    TIMEOUT = "20"
  }
}

resource "kubernetes_secret" "mysql_secret" {
  metadata {
    name = "mysql-secret"
  }

  data = {
    mysql-db = "db"

    mysql-pass = "password"

    mysql-root-pass = "password"

    mysql-user = "admin"
  }

  type = "Opaque"
}

resource "kubernetes_deployment" "mysql_deployment" {
  metadata {
    name = "mysql-deployment"

    labels = {
      app = "mysql"
    }
  }

  spec {
    replicas = 1

    selector {
      match_labels = {
        app = "mysql"
      }
    }

    template {
      metadata {
        labels = {
          app = "mysql"
        }
      }

      spec {
        container {
          name  = "mysql"
          image = "mysql"

          port {
            container_port = 3306
          }

          env {
            name = "MYSQL_ROOT_PASSWORD"

            value_from {
              secret_key_ref {
                name = "mysql-secret"
                key  = "mysql-root-pass"
              }
            }
          }

          env {
            name = "MYSQL_USER"

            value_from {
              secret_key_ref {
                name = "mysql-secret"
                key  = "mysql-user"
              }
            }
          }

          env {
            name = "MYSQL_DATABASE"

            value_from {
              secret_key_ref {
                name = "mysql-secret"
                key  = "mysql-db"
              }
            }
          }

          env {
            name = "MYSQL_PASSWORD"

            value_from {
              secret_key_ref {
                name = "mysql-secret"
                key  = "mysql-pass"
              }
            }
          }
        }
      }
    }
  }
}

resource "kubernetes_service" "mysql_service" {
  metadata {
    name = "mysql-service"
  }

  spec {
    port {
      protocol    = "TCP"
      port        = 3306
      target_port = "3306"
    }

    selector = {
      app = "mysql"
    }
  }
}

resource "kubernetes_deployment" "server_deployment" {
  metadata {
    name = "server-deployment"

    labels = {
      app = "server"
    }
  }

  spec {
    replicas = 2

    selector {
      match_labels = {
        app = "server"
      }
    }

    template {
      metadata {
        labels = {
          app = "server"
        }
      }

      spec {
        container {
          name  = "server"
          image = "ericmarcantonio/pic-it-server"

          port {
            container_port = 8000
          }

          env {
            name = "USER"

            value_from {
              secret_key_ref {
                name = "mysql-secret"
                key  = "mysql-user"
              }
            }
          }

          env {
            name = "PASS"

            value_from {
              secret_key_ref {
                name = "mysql-secret"
                key  = "mysql-pass"
              }
            }
          }

          env {
            name = "HOST"

            value_from {
              config_map_key_ref {
                name = "pic-it-configmap"
                key  = "HOST"
              }
            }
          }

          env {
            name = "PORT"

            value_from {
              config_map_key_ref {
                name = "pic-it-configmap"
                key  = "PORT"
              }
            }
          }

          env {
            name = "DB"

            value_from {
              secret_key_ref {
                name = "mysql-secret"
                key  = "mysql-db"
              }
            }
          }

          env {
            name = "TIMEOUT"

            value_from {
              config_map_key_ref {
                name = "pic-it-configmap"
                key  = "TIMEOUT"
              }
            }
          }

          resources {
            limits = {
              cpu    = "500m"
              memory = "1Gi"
            }

            requests = {
              cpu    = "200m"
              memory = "500m"
            }
          }

          image_pull_policy = "Always"
        }
      }
    }
  }
}

resource "kubernetes_service" "server_external_service" {
  metadata {
    name = "server-external-service"
  }

  spec {
    port {
      protocol    = "TCP"
      port        = 8000
      target_port = "8000"
      node_port   = 30000
    }

    selector = {
      app = "server"
    }

    type = "LoadBalancer"
  }
}

resource "kubernetes_horizontal_pod_autoscaler" "server_scaler" {
  metadata {
    name = "server-scaler"
  }

  spec {
    scale_target_ref {
      kind = "Deployment"
      name = "server-deployment"
    }

    min_replicas = 2
    max_replicas = 30

    metric {
      type = "Resource"

      resource {
        name = "memory"

        target {
          type = "AverageValue"
        }
      }
    }

    metric {
      type = "Resource"

      resource {
        name = "cpu"

        target {
          type                = "Utilization"
          average_utilization = 10
        }
      }
    }
  }
}

resource "kubernetes_service_account" "admin_user" {
  metadata {
    name      = "admin-user"
    namespace = "kube-system"
  }
}

resource "kubernetes_cluster_role_binding" "admin_user" {
  metadata {
    name = "admin-user"
  }

  subject {
    kind      = "ServiceAccount"
    name      = "admin-user"
    namespace = "kube-system"
  }

  role_ref {
    api_group = "rbac.authorization.k8s.io"
    kind      = "ClusterRole"
    name      = "cluster-admin"
  }
}

