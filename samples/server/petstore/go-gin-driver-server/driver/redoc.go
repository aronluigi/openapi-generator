/*
 * OpenAPI Petstore
 *
 * This is a sample server Petstore server. For this sample, you can use the api key `special-key` to test the authorization filters.
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package petstoreserver

const ReDocHTML string = `PCFET0NUWVBFIGh0bWw+CjxodG1sPgogICAgPGhlYWQ+CiAgICAgICAgPHRpdGxlPk9wZW5BUEkgUGV0c3RvcmU8L3RpdGxlPgogICAgICAgIDwhLS0gbmVlZGVkIGZvciBhZGFwdGl2ZSBkZXNpZ24gLS0+CiAgICAgICAgPG1ldGEgY2hhcnNldD0idXRmLTgiLz4KICAgICAgICA8bWV0YSBuYW1lPSJ2aWV3cG9ydCIgY29udGVudD0id2lkdGg9ZGV2aWNlLXdpZHRoLCBpbml0aWFsLXNjYWxlPTEiPgogICAgICAgIDxsaW5rIGhyZWY9Imh0dHBzOi8vZm9udHMuZ29vZ2xlYXBpcy5jb20vY3NzP2ZhbWlseT1Nb250c2VycmF0OjMwMCw0MDAsNzAwfFJvYm90bzozMDAsNDAwLDcwMCIgcmVsPSJzdHlsZXNoZWV0Ij4KCiAgICAgICAgPCEtLQogICAgICAgIFJlRG9jIGRvZXNuJ3QgY2hhbmdlIG91dGVyIHBhZ2Ugc3R5bGVzCiAgICAgICAgLS0+CiAgICAgICAgPHN0eWxlPgogICAgICAgIGJvZHkgewogICAgICAgICAgICBtYXJnaW46IDA7CiAgICAgICAgICAgIHBhZGRpbmc6IDA7CiAgICAgICAgfQogICAgICAgIDwvc3R5bGU+CiAgICA8L2hlYWQ+CiAgICA8Ym9keT4KICAgICAgICA8cmVkb2Mgc3BlYy11cmw9Jy92Mi9kb2MvYXBpX3NwZWMueWFtbCc+PC9yZWRvYz4KICAgICAgICA8c2NyaXB0IHNyYz0iaHR0cHM6Ly9jZG4uanNkZWxpdnIubmV0L25wbS9yZWRvY0BuZXh0L2J1bmRsZXMvcmVkb2Muc3RhbmRhbG9uZS5qcyI+IDwvc2NyaXB0PgogICAgPC9ib2R5Pgo8L2h0bWw+`

const ReDocAPISpec string = `b3BlbmFwaTogMy4wLjEKaW5mbzoKICBkZXNjcmlwdGlvbjogVGhpcyBpcyBhIHNhbXBsZSBzZXJ2ZXIgUGV0c3RvcmUgc2VydmVyLiBGb3IgdGhpcyBzYW1wbGUsIHlvdSBjYW4gdXNlCiAgICB0aGUgYXBpIGtleSBgc3BlY2lhbC1rZXlgIHRvIHRlc3QgdGhlIGF1dGhvcml6YXRpb24gZmlsdGVycy4KICBsaWNlbnNlOgogICAgbmFtZTogQXBhY2hlLTIuMAogICAgdXJsOiBodHRwOi8vd3d3LmFwYWNoZS5vcmcvbGljZW5zZXMvTElDRU5TRS0yLjAuaHRtbAogIHRpdGxlOiBPcGVuQVBJIFBldHN0b3JlCiAgdmVyc2lvbjogMS4wLjAKc2VydmVyczoKLSB1cmw6IGh0dHA6Ly9wZXRzdG9yZS5zd2FnZ2VyLmlvL3YyCnRhZ3M6Ci0gZGVzY3JpcHRpb246IEV2ZXJ5dGhpbmcgYWJvdXQgeW91ciBQZXRzCiAgbmFtZTogcGV0Ci0gZGVzY3JpcHRpb246IEFjY2VzcyB0byBQZXRzdG9yZSBvcmRlcnMKICBuYW1lOiBzdG9yZQotIGRlc2NyaXB0aW9uOiBPcGVyYXRpb25zIGFib3V0IHVzZXIKICBuYW1lOiB1c2VyCnBhdGhzOgogIC9wZXQ6CiAgICBwb3N0OgogICAgICBvcGVyYXRpb25JZDogYWRkUGV0CiAgICAgIHJlcXVlc3RCb2R5OgogICAgICAgIGNvbnRlbnQ6CiAgICAgICAgICBhcHBsaWNhdGlvbi9qc29uOgogICAgICAgICAgICBzY2hlbWE6CiAgICAgICAgICAgICAgJHJlZjogJyMvY29tcG9uZW50cy9zY2hlbWFzL1BldCcKICAgICAgICAgIGFwcGxpY2F0aW9uL3htbDoKICAgICAgICAgICAgc2NoZW1hOgogICAgICAgICAgICAgICRyZWY6ICcjL2NvbXBvbmVudHMvc2NoZW1hcy9QZXQnCiAgICAgICAgZGVzY3JpcHRpb246IFBldCBvYmplY3QgdGhhdCBuZWVkcyB0byBiZSBhZGRlZCB0byB0aGUgc3RvcmUKICAgICAgICByZXF1aXJlZDogdHJ1ZQogICAgICByZXNwb25zZXM6CiAgICAgICAgNDA1OgogICAgICAgICAgY29udGVudDoge30KICAgICAgICAgIGRlc2NyaXB0aW9uOiBJbnZhbGlkIGlucHV0CiAgICAgIHNlY3VyaXR5OgogICAgICAtIHBldHN0b3JlX2F1dGg6CiAgICAgICAgLSB3cml0ZTpwZXRzCiAgICAgICAgLSByZWFkOnBldHMKICAgICAgc3VtbWFyeTogQWRkIGEgbmV3IHBldCB0byB0aGUgc3RvcmUKICAgICAgdGFnczoKICAgICAgLSBwZXQKICAgICAgeC1jb2RlZ2VuLXJlcXVlc3QtYm9keS1uYW1lOiBib2R5CiAgICBwdXQ6CiAgICAgIG9wZXJhdGlvbklkOiB1cGRhdGVQZXQKICAgICAgcmVxdWVzdEJvZHk6CiAgICAgICAgY29udGVudDoKICAgICAgICAgIGFwcGxpY2F0aW9uL2pzb246CiAgICAgICAgICAgIHNjaGVtYToKICAgICAgICAgICAgICAkcmVmOiAnIy9jb21wb25lbnRzL3NjaGVtYXMvUGV0JwogICAgICAgICAgYXBwbGljYXRpb24veG1sOgogICAgICAgICAgICBzY2hlbWE6CiAgICAgICAgICAgICAgJHJlZjogJyMvY29tcG9uZW50cy9zY2hlbWFzL1BldCcKICAgICAgICBkZXNjcmlwdGlvbjogUGV0IG9iamVjdCB0aGF0IG5lZWRzIHRvIGJlIGFkZGVkIHRvIHRoZSBzdG9yZQogICAgICAgIHJlcXVpcmVkOiB0cnVlCiAgICAgIHJlc3BvbnNlczoKICAgICAgICA0MDA6CiAgICAgICAgICBjb250ZW50OiB7fQogICAgICAgICAgZGVzY3JpcHRpb246IEludmFsaWQgSUQgc3VwcGxpZWQKICAgICAgICA0MDQ6CiAgICAgICAgICBjb250ZW50OiB7fQogICAgICAgICAgZGVzY3JpcHRpb246IFBldCBub3QgZm91bmQKICAgICAgICA0MDU6CiAgICAgICAgICBjb250ZW50OiB7fQogICAgICAgICAgZGVzY3JpcHRpb246IFZhbGlkYXRpb24gZXhjZXB0aW9uCiAgICAgIHNlY3VyaXR5OgogICAgICAtIHBldHN0b3JlX2F1dGg6CiAgICAgICAgLSB3cml0ZTpwZXRzCiAgICAgICAgLSByZWFkOnBldHMKICAgICAgc3VtbWFyeTogVXBkYXRlIGFuIGV4aXN0aW5nIHBldAogICAgICB0YWdzOgogICAgICAtIHBldAogICAgICB4LWNvZGVnZW4tcmVxdWVzdC1ib2R5LW5hbWU6IGJvZHkKICAvcGV0L2ZpbmRCeVN0YXR1czoKICAgIGdldDoKICAgICAgZGVzY3JpcHRpb246IE11bHRpcGxlIHN0YXR1cyB2YWx1ZXMgY2FuIGJlIHByb3ZpZGVkIHdpdGggY29tbWEgc2VwYXJhdGVkIHN0cmluZ3MKICAgICAgb3BlcmF0aW9uSWQ6IGZpbmRQZXRzQnlTdGF0dXMKICAgICAgcGFyYW1ldGVyczoKICAgICAgLSBkZXNjcmlwdGlvbjogU3RhdHVzIHZhbHVlcyB0aGF0IG5lZWQgdG8gYmUgY29uc2lkZXJlZCBmb3IgZmlsdGVyCiAgICAgICAgZXhwbG9kZTogZmFsc2UKICAgICAgICBpbjogcXVlcnkKICAgICAgICBuYW1lOiBzdGF0dXMKICAgICAgICByZXF1aXJlZDogdHJ1ZQogICAgICAgIHNjaGVtYToKICAgICAgICAgIGl0ZW1zOgogICAgICAgICAgICBkZWZhdWx0OiBhdmFpbGFibGUKICAgICAgICAgICAgZW51bToKICAgICAgICAgICAgLSBhdmFpbGFibGUKICAgICAgICAgICAgLSBwZW5kaW5nCiAgICAgICAgICAgIC0gc29sZAogICAgICAgICAgICB0eXBlOiBzdHJpbmcKICAgICAgICAgIHR5cGU6IGFycmF5CiAgICAgICAgc3R5bGU6IGZvcm0KICAgICAgcmVzcG9uc2VzOgogICAgICAgIDIwMDoKICAgICAgICAgIGNvbnRlbnQ6CiAgICAgICAgICAgIGFwcGxpY2F0aW9uL3htbDoKICAgICAgICAgICAgICBzY2hlbWE6CiAgICAgICAgICAgICAgICBpdGVtczoKICAgICAgICAgICAgICAgICAgJHJlZjogJyMvY29tcG9uZW50cy9zY2hlbWFzL1BldCcKICAgICAgICAgICAgICAgIHR5cGU6IGFycmF5CiAgICAgICAgICAgIGFwcGxpY2F0aW9uL2pzb246CiAgICAgICAgICAgICAgc2NoZW1hOgogICAgICAgICAgICAgICAgaXRlbXM6CiAgICAgICAgICAgICAgICAgICRyZWY6ICcjL2NvbXBvbmVudHMvc2NoZW1hcy9QZXQnCiAgICAgICAgICAgICAgICB0eXBlOiBhcnJheQogICAgICAgICAgZGVzY3JpcHRpb246IHN1Y2Nlc3NmdWwgb3BlcmF0aW9uCiAgICAgICAgNDAwOgogICAgICAgICAgY29udGVudDoge30KICAgICAgICAgIGRlc2NyaXB0aW9uOiBJbnZhbGlkIHN0YXR1cyB2YWx1ZQogICAgICBzZWN1cml0eToKICAgICAgLSBwZXRzdG9yZV9hdXRoOgogICAgICAgIC0gd3JpdGU6cGV0cwogICAgICAgIC0gcmVhZDpwZXRzCiAgICAgIHN1bW1hcnk6IEZpbmRzIFBldHMgYnkgc3RhdHVzCiAgICAgIHRhZ3M6CiAgICAgIC0gcGV0CiAgL3BldC9maW5kQnlUYWdzOgogICAgZ2V0OgogICAgICBkZXByZWNhdGVkOiB0cnVlCiAgICAgIGRlc2NyaXB0aW9uOiBNdWx0aXBsZSB0YWdzIGNhbiBiZSBwcm92aWRlZCB3aXRoIGNvbW1hIHNlcGFyYXRlZCBzdHJpbmdzLiBVc2UKICAgICAgICB0YWcxLCB0YWcyLCB0YWczIGZvciB0ZXN0aW5nLgogICAgICBvcGVyYXRpb25JZDogZmluZFBldHNCeVRhZ3MKICAgICAgcGFyYW1ldGVyczoKICAgICAgLSBkZXNjcmlwdGlvbjogVGFncyB0byBmaWx0ZXIgYnkKICAgICAgICBleHBsb2RlOiBmYWxzZQogICAgICAgIGluOiBxdWVyeQogICAgICAgIG5hbWU6IHRhZ3MKICAgICAgICByZXF1aXJlZDogdHJ1ZQogICAgICAgIHNjaGVtYToKICAgICAgICAgIGl0ZW1zOgogICAgICAgICAgICB0eXBlOiBzdHJpbmcKICAgICAgICAgIHR5cGU6IGFycmF5CiAgICAgICAgc3R5bGU6IGZvcm0KICAgICAgcmVzcG9uc2VzOgogICAgICAgIDIwMDoKICAgICAgICAgIGNvbnRlbnQ6CiAgICAgICAgICAgIGFwcGxpY2F0aW9uL3htbDoKICAgICAgICAgICAgICBzY2hlbWE6CiAgICAgICAgICAgICAgICBpdGVtczoKICAgICAgICAgICAgICAgICAgJHJlZjogJyMvY29tcG9uZW50cy9zY2hlbWFzL1BldCcKICAgICAgICAgICAgICAgIHR5cGU6IGFycmF5CiAgICAgICAgICAgIGFwcGxpY2F0aW9uL2pzb246CiAgICAgICAgICAgICAgc2NoZW1hOgogICAgICAgICAgICAgICAgaXRlbXM6CiAgICAgICAgICAgICAgICAgICRyZWY6ICcjL2NvbXBvbmVudHMvc2NoZW1hcy9QZXQnCiAgICAgICAgICAgICAgICB0eXBlOiBhcnJheQogICAgICAgICAgZGVzY3JpcHRpb246IHN1Y2Nlc3NmdWwgb3BlcmF0aW9uCiAgICAgICAgNDAwOgogICAgICAgICAgY29udGVudDoge30KICAgICAgICAgIGRlc2NyaXB0aW9uOiBJbnZhbGlkIHRhZyB2YWx1ZQogICAgICBzZWN1cml0eToKICAgICAgLSBwZXRzdG9yZV9hdXRoOgogICAgICAgIC0gd3JpdGU6cGV0cwogICAgICAgIC0gcmVhZDpwZXRzCiAgICAgIHN1bW1hcnk6IEZpbmRzIFBldHMgYnkgdGFncwogICAgICB0YWdzOgogICAgICAtIHBldAogIC9wZXQve3BldElkfToKICAgIGRlbGV0ZToKICAgICAgb3BlcmF0aW9uSWQ6IGRlbGV0ZVBldAogICAgICBwYXJhbWV0ZXJzOgogICAgICAtIGluOiBoZWFkZXIKICAgICAgICBuYW1lOiBhcGlfa2V5CiAgICAgICAgc2NoZW1hOgogICAgICAgICAgdHlwZTogc3RyaW5nCiAgICAgIC0gZGVzY3JpcHRpb246IFBldCBpZCB0byBkZWxldGUKICAgICAgICBpbjogcGF0aAogICAgICAgIG5hbWU6IHBldElkCiAgICAgICAgcmVxdWlyZWQ6IHRydWUKICAgICAgICBzY2hlbWE6CiAgICAgICAgICBmb3JtYXQ6IGludDY0CiAgICAgICAgICB0eXBlOiBpbnRlZ2VyCiAgICAgIHJlc3BvbnNlczoKICAgICAgICA0MDA6CiAgICAgICAgICBjb250ZW50OiB7fQogICAgICAgICAgZGVzY3JpcHRpb246IEludmFsaWQgcGV0IHZhbHVlCiAgICAgIHNlY3VyaXR5OgogICAgICAtIHBldHN0b3JlX2F1dGg6CiAgICAgICAgLSB3cml0ZTpwZXRzCiAgICAgICAgLSByZWFkOnBldHMKICAgICAgc3VtbWFyeTogRGVsZXRlcyBhIHBldAogICAgICB0YWdzOgogICAgICAtIHBldAogICAgZ2V0OgogICAgICBkZXNjcmlwdGlvbjogUmV0dXJucyBhIHNpbmdsZSBwZXQKICAgICAgb3BlcmF0aW9uSWQ6IGdldFBldEJ5SWQKICAgICAgcGFyYW1ldGVyczoKICAgICAgLSBkZXNjcmlwdGlvbjogSUQgb2YgcGV0IHRvIHJldHVybgogICAgICAgIGluOiBwYXRoCiAgICAgICAgbmFtZTogcGV0SWQKICAgICAgICByZXF1aXJlZDogdHJ1ZQogICAgICAgIHNjaGVtYToKICAgICAgICAgIGZvcm1hdDogaW50NjQKICAgICAgICAgIHR5cGU6IGludGVnZXIKICAgICAgcmVzcG9uc2VzOgogICAgICAgIDIwMDoKICAgICAgICAgIGNvbnRlbnQ6CiAgICAgICAgICAgIGFwcGxpY2F0aW9uL3htbDoKICAgICAgICAgICAgICBzY2hlbWE6CiAgICAgICAgICAgICAgICAkcmVmOiAnIy9jb21wb25lbnRzL3NjaGVtYXMvUGV0JwogICAgICAgICAgICBhcHBsaWNhdGlvbi9qc29uOgogICAgICAgICAgICAgIHNjaGVtYToKICAgICAgICAgICAgICAgICRyZWY6ICcjL2NvbXBvbmVudHMvc2NoZW1hcy9QZXQnCiAgICAgICAgICBkZXNjcmlwdGlvbjogc3VjY2Vzc2Z1bCBvcGVyYXRpb24KICAgICAgICA0MDA6CiAgICAgICAgICBjb250ZW50OiB7fQogICAgICAgICAgZGVzY3JpcHRpb246IEludmFsaWQgSUQgc3VwcGxpZWQKICAgICAgICA0MDQ6CiAgICAgICAgICBjb250ZW50OiB7fQogICAgICAgICAgZGVzY3JpcHRpb246IFBldCBub3QgZm91bmQKICAgICAgc2VjdXJpdHk6CiAgICAgIC0gYXBpX2tleTogW10KICAgICAgc3VtbWFyeTogRmluZCBwZXQgYnkgSUQKICAgICAgdGFnczoKICAgICAgLSBwZXQKICAgIHBvc3Q6CiAgICAgIG9wZXJhdGlvbklkOiB1cGRhdGVQZXRXaXRoRm9ybQogICAgICBwYXJhbWV0ZXJzOgogICAgICAtIGRlc2NyaXB0aW9uOiBJRCBvZiBwZXQgdGhhdCBuZWVkcyB0byBiZSB1cGRhdGVkCiAgICAgICAgaW46IHBhdGgKICAgICAgICBuYW1lOiBwZXRJZAogICAgICAgIHJlcXVpcmVkOiB0cnVlCiAgICAgICAgc2NoZW1hOgogICAgICAgICAgZm9ybWF0OiBpbnQ2NAogICAgICAgICAgdHlwZTogaW50ZWdlcgogICAgICByZXF1ZXN0Qm9keToKICAgICAgICBjb250ZW50OgogICAgICAgICAgYXBwbGljYXRpb24veC13d3ctZm9ybS11cmxlbmNvZGVkOgogICAgICAgICAgICBzY2hlbWE6CiAgICAgICAgICAgICAgcHJvcGVydGllczoKICAgICAgICAgICAgICAgIG5hbWU6CiAgICAgICAgICAgICAgICAgIGRlc2NyaXB0aW9uOiBVcGRhdGVkIG5hbWUgb2YgdGhlIHBldAogICAgICAgICAgICAgICAgICB0eXBlOiBzdHJpbmcKICAgICAgICAgICAgICAgIHN0YXR1czoKICAgICAgICAgICAgICAgICAgZGVzY3JpcHRpb246IFVwZGF0ZWQgc3RhdHVzIG9mIHRoZSBwZXQKICAgICAgICAgICAgICAgICAgdHlwZTogc3RyaW5nCiAgICAgIHJlc3BvbnNlczoKICAgICAgICA0MDU6CiAgICAgICAgICBjb250ZW50OiB7fQogICAgICAgICAgZGVzY3JpcHRpb246IEludmFsaWQgaW5wdXQKICAgICAgc2VjdXJpdHk6CiAgICAgIC0gcGV0c3RvcmVfYXV0aDoKICAgICAgICAtIHdyaXRlOnBldHMKICAgICAgICAtIHJlYWQ6cGV0cwogICAgICBzdW1tYXJ5OiBVcGRhdGVzIGEgcGV0IGluIHRoZSBzdG9yZSB3aXRoIGZvcm0gZGF0YQogICAgICB0YWdzOgogICAgICAtIHBldAogIC9wZXQve3BldElkfS91cGxvYWRJbWFnZToKICAgIHBvc3Q6CiAgICAgIG9wZXJhdGlvbklkOiB1cGxvYWRGaWxlCiAgICAgIHBhcmFtZXRlcnM6CiAgICAgIC0gZGVzY3JpcHRpb246IElEIG9mIHBldCB0byB1cGRhdGUKICAgICAgICBpbjogcGF0aAogICAgICAgIG5hbWU6IHBldElkCiAgICAgICAgcmVxdWlyZWQ6IHRydWUKICAgICAgICBzY2hlbWE6CiAgICAgICAgICBmb3JtYXQ6IGludDY0CiAgICAgICAgICB0eXBlOiBpbnRlZ2VyCiAgICAgIHJlcXVlc3RCb2R5OgogICAgICAgIGNvbnRlbnQ6CiAgICAgICAgICBtdWx0aXBhcnQvZm9ybS1kYXRhOgogICAgICAgICAgICBzY2hlbWE6CiAgICAgICAgICAgICAgcHJvcGVydGllczoKICAgICAgICAgICAgICAgIGFkZGl0aW9uYWxNZXRhZGF0YToKICAgICAgICAgICAgICAgICAgZGVzY3JpcHRpb246IEFkZGl0aW9uYWwgZGF0YSB0byBwYXNzIHRvIHNlcnZlcgogICAgICAgICAgICAgICAgICB0eXBlOiBzdHJpbmcKICAgICAgICAgICAgICAgIGZpbGU6CiAgICAgICAgICAgICAgICAgIGRlc2NyaXB0aW9uOiBmaWxlIHRvIHVwbG9hZAogICAgICAgICAgICAgICAgICBmb3JtYXQ6IGJpbmFyeQogICAgICAgICAgICAgICAgICB0eXBlOiBzdHJpbmcKICAgICAgcmVzcG9uc2VzOgogICAgICAgIDIwMDoKICAgICAgICAgIGNvbnRlbnQ6CiAgICAgICAgICAgIGFwcGxpY2F0aW9uL2pzb246CiAgICAgICAgICAgICAgc2NoZW1hOgogICAgICAgICAgICAgICAgJHJlZjogJyMvY29tcG9uZW50cy9zY2hlbWFzL0FwaVJlc3BvbnNlJwogICAgICAgICAgZGVzY3JpcHRpb246IHN1Y2Nlc3NmdWwgb3BlcmF0aW9uCiAgICAgIHNlY3VyaXR5OgogICAgICAtIHBldHN0b3JlX2F1dGg6CiAgICAgICAgLSB3cml0ZTpwZXRzCiAgICAgICAgLSByZWFkOnBldHMKICAgICAgc3VtbWFyeTogdXBsb2FkcyBhbiBpbWFnZQogICAgICB0YWdzOgogICAgICAtIHBldAogIC9zdG9yZS9pbnZlbnRvcnk6CiAgICBnZXQ6CiAgICAgIGRlc2NyaXB0aW9uOiBSZXR1cm5zIGEgbWFwIG9mIHN0YXR1cyBjb2RlcyB0byBxdWFudGl0aWVzCiAgICAgIG9wZXJhdGlvbklkOiBnZXRJbnZlbnRvcnkKICAgICAgcmVzcG9uc2VzOgogICAgICAgIDIwMDoKICAgICAgICAgIGNvbnRlbnQ6CiAgICAgICAgICAgIGFwcGxpY2F0aW9uL2pzb246CiAgICAgICAgICAgICAgc2NoZW1hOgogICAgICAgICAgICAgICAgYWRkaXRpb25hbFByb3BlcnRpZXM6CiAgICAgICAgICAgICAgICAgIGZvcm1hdDogaW50MzIKICAgICAgICAgICAgICAgICAgdHlwZTogaW50ZWdlcgogICAgICAgICAgICAgICAgdHlwZTogb2JqZWN0CiAgICAgICAgICBkZXNjcmlwdGlvbjogc3VjY2Vzc2Z1bCBvcGVyYXRpb24KICAgICAgc2VjdXJpdHk6CiAgICAgIC0gYXBpX2tleTogW10KICAgICAgc3VtbWFyeTogUmV0dXJucyBwZXQgaW52ZW50b3JpZXMgYnkgc3RhdHVzCiAgICAgIHRhZ3M6CiAgICAgIC0gc3RvcmUKICAvc3RvcmUvb3JkZXI6CiAgICBwb3N0OgogICAgICBvcGVyYXRpb25JZDogcGxhY2VPcmRlcgogICAgICByZXF1ZXN0Qm9keToKICAgICAgICBjb250ZW50OgogICAgICAgICAgJyovKic6CiAgICAgICAgICAgIHNjaGVtYToKICAgICAgICAgICAgICAkcmVmOiAnIy9jb21wb25lbnRzL3NjaGVtYXMvT3JkZXInCiAgICAgICAgZGVzY3JpcHRpb246IG9yZGVyIHBsYWNlZCBmb3IgcHVyY2hhc2luZyB0aGUgcGV0CiAgICAgICAgcmVxdWlyZWQ6IHRydWUKICAgICAgcmVzcG9uc2VzOgogICAgICAgIDIwMDoKICAgICAgICAgIGNvbnRlbnQ6CiAgICAgICAgICAgIGFwcGxpY2F0aW9uL3htbDoKICAgICAgICAgICAgICBzY2hlbWE6CiAgICAgICAgICAgICAgICAkcmVmOiAnIy9jb21wb25lbnRzL3NjaGVtYXMvT3JkZXInCiAgICAgICAgICAgIGFwcGxpY2F0aW9uL2pzb246CiAgICAgICAgICAgICAgc2NoZW1hOgogICAgICAgICAgICAgICAgJHJlZjogJyMvY29tcG9uZW50cy9zY2hlbWFzL09yZGVyJwogICAgICAgICAgZGVzY3JpcHRpb246IHN1Y2Nlc3NmdWwgb3BlcmF0aW9uCiAgICAgICAgNDAwOgogICAgICAgICAgY29udGVudDoge30KICAgICAgICAgIGRlc2NyaXB0aW9uOiBJbnZhbGlkIE9yZGVyCiAgICAgIHN1bW1hcnk6IFBsYWNlIGFuIG9yZGVyIGZvciBhIHBldAogICAgICB0YWdzOgogICAgICAtIHN0b3JlCiAgICAgIHgtY29kZWdlbi1yZXF1ZXN0LWJvZHktbmFtZTogYm9keQogIC9zdG9yZS9vcmRlci97b3JkZXJJZH06CiAgICBkZWxldGU6CiAgICAgIGRlc2NyaXB0aW9uOiBGb3IgdmFsaWQgcmVzcG9uc2UgdHJ5IGludGVnZXIgSURzIHdpdGggdmFsdWUgPCAxMDAwLiBBbnl0aGluZwogICAgICAgIGFib3ZlIDEwMDAgb3Igbm9uaW50ZWdlcnMgd2lsbCBnZW5lcmF0ZSBBUEkgZXJyb3JzCiAgICAgIG9wZXJhdGlvbklkOiBkZWxldGVPcmRlcgogICAgICBwYXJhbWV0ZXJzOgogICAgICAtIGRlc2NyaXB0aW9uOiBJRCBvZiB0aGUgb3JkZXIgdGhhdCBuZWVkcyB0byBiZSBkZWxldGVkCiAgICAgICAgaW46IHBhdGgKICAgICAgICBuYW1lOiBvcmRlcklkCiAgICAgICAgcmVxdWlyZWQ6IHRydWUKICAgICAgICBzY2hlbWE6CiAgICAgICAgICB0eXBlOiBzdHJpbmcKICAgICAgcmVzcG9uc2VzOgogICAgICAgIDQwMDoKICAgICAgICAgIGNvbnRlbnQ6IHt9CiAgICAgICAgICBkZXNjcmlwdGlvbjogSW52YWxpZCBJRCBzdXBwbGllZAogICAgICAgIDQwNDoKICAgICAgICAgIGNvbnRlbnQ6IHt9CiAgICAgICAgICBkZXNjcmlwdGlvbjogT3JkZXIgbm90IGZvdW5kCiAgICAgIHN1bW1hcnk6IERlbGV0ZSBwdXJjaGFzZSBvcmRlciBieSBJRAogICAgICB0YWdzOgogICAgICAtIHN0b3JlCiAgICBnZXQ6CiAgICAgIGRlc2NyaXB0aW9uOiBGb3IgdmFsaWQgcmVzcG9uc2UgdHJ5IGludGVnZXIgSURzIHdpdGggdmFsdWUgPD0gNSBvciA+IDEwLiBPdGhlcgogICAgICAgIHZhbHVlcyB3aWxsIGdlbmVyYXRlZCBleGNlcHRpb25zCiAgICAgIG9wZXJhdGlvbklkOiBnZXRPcmRlckJ5SWQKICAgICAgcGFyYW1ldGVyczoKICAgICAgLSBkZXNjcmlwdGlvbjogSUQgb2YgcGV0IHRoYXQgbmVlZHMgdG8gYmUgZmV0Y2hlZAogICAgICAgIGluOiBwYXRoCiAgICAgICAgbmFtZTogb3JkZXJJZAogICAgICAgIHJlcXVpcmVkOiB0cnVlCiAgICAgICAgc2NoZW1hOgogICAgICAgICAgZm9ybWF0OiBpbnQ2NAogICAgICAgICAgbWF4aW11bTogNQogICAgICAgICAgbWluaW11bTogMQogICAgICAgICAgdHlwZTogaW50ZWdlcgogICAgICByZXNwb25zZXM6CiAgICAgICAgMjAwOgogICAgICAgICAgY29udGVudDoKICAgICAgICAgICAgYXBwbGljYXRpb24veG1sOgogICAgICAgICAgICAgIHNjaGVtYToKICAgICAgICAgICAgICAgICRyZWY6ICcjL2NvbXBvbmVudHMvc2NoZW1hcy9PcmRlcicKICAgICAgICAgICAgYXBwbGljYXRpb24vanNvbjoKICAgICAgICAgICAgICBzY2hlbWE6CiAgICAgICAgICAgICAgICAkcmVmOiAnIy9jb21wb25lbnRzL3NjaGVtYXMvT3JkZXInCiAgICAgICAgICBkZXNjcmlwdGlvbjogc3VjY2Vzc2Z1bCBvcGVyYXRpb24KICAgICAgICA0MDA6CiAgICAgICAgICBjb250ZW50OiB7fQogICAgICAgICAgZGVzY3JpcHRpb246IEludmFsaWQgSUQgc3VwcGxpZWQKICAgICAgICA0MDQ6CiAgICAgICAgICBjb250ZW50OiB7fQogICAgICAgICAgZGVzY3JpcHRpb246IE9yZGVyIG5vdCBmb3VuZAogICAgICBzdW1tYXJ5OiBGaW5kIHB1cmNoYXNlIG9yZGVyIGJ5IElECiAgICAgIHRhZ3M6CiAgICAgIC0gc3RvcmUKICAvdXNlcjoKICAgIHBvc3Q6CiAgICAgIGRlc2NyaXB0aW9uOiBUaGlzIGNhbiBvbmx5IGJlIGRvbmUgYnkgdGhlIGxvZ2dlZCBpbiB1c2VyLgogICAgICBvcGVyYXRpb25JZDogY3JlYXRlVXNlcgogICAgICByZXF1ZXN0Qm9keToKICAgICAgICBjb250ZW50OgogICAgICAgICAgJyovKic6CiAgICAgICAgICAgIHNjaGVtYToKICAgICAgICAgICAgICAkcmVmOiAnIy9jb21wb25lbnRzL3NjaGVtYXMvVXNlcicKICAgICAgICBkZXNjcmlwdGlvbjogQ3JlYXRlZCB1c2VyIG9iamVjdAogICAgICAgIHJlcXVpcmVkOiB0cnVlCiAgICAgIHJlc3BvbnNlczoKICAgICAgICBkZWZhdWx0OgogICAgICAgICAgY29udGVudDoge30KICAgICAgICAgIGRlc2NyaXB0aW9uOiBzdWNjZXNzZnVsIG9wZXJhdGlvbgogICAgICBzdW1tYXJ5OiBDcmVhdGUgdXNlcgogICAgICB0YWdzOgogICAgICAtIHVzZXIKICAgICAgeC1jb2RlZ2VuLXJlcXVlc3QtYm9keS1uYW1lOiBib2R5CiAgL3VzZXIvY3JlYXRlV2l0aEFycmF5OgogICAgcG9zdDoKICAgICAgb3BlcmF0aW9uSWQ6IGNyZWF0ZVVzZXJzV2l0aEFycmF5SW5wdXQKICAgICAgcmVxdWVzdEJvZHk6CiAgICAgICAgY29udGVudDoKICAgICAgICAgICcqLyonOgogICAgICAgICAgICBzY2hlbWE6CiAgICAgICAgICAgICAgaXRlbXM6CiAgICAgICAgICAgICAgICAkcmVmOiAnIy9jb21wb25lbnRzL3NjaGVtYXMvVXNlcicKICAgICAgICAgICAgICB0eXBlOiBhcnJheQogICAgICAgIGRlc2NyaXB0aW9uOiBMaXN0IG9mIHVzZXIgb2JqZWN0CiAgICAgICAgcmVxdWlyZWQ6IHRydWUKICAgICAgcmVzcG9uc2VzOgogICAgICAgIGRlZmF1bHQ6CiAgICAgICAgICBjb250ZW50OiB7fQogICAgICAgICAgZGVzY3JpcHRpb246IHN1Y2Nlc3NmdWwgb3BlcmF0aW9uCiAgICAgIHN1bW1hcnk6IENyZWF0ZXMgbGlzdCBvZiB1c2VycyB3aXRoIGdpdmVuIGlucHV0IGFycmF5CiAgICAgIHRhZ3M6CiAgICAgIC0gdXNlcgogICAgICB4LWNvZGVnZW4tcmVxdWVzdC1ib2R5LW5hbWU6IGJvZHkKICAvdXNlci9jcmVhdGVXaXRoTGlzdDoKICAgIHBvc3Q6CiAgICAgIG9wZXJhdGlvbklkOiBjcmVhdGVVc2Vyc1dpdGhMaXN0SW5wdXQKICAgICAgcmVxdWVzdEJvZHk6CiAgICAgICAgY29udGVudDoKICAgICAgICAgICcqLyonOgogICAgICAgICAgICBzY2hlbWE6CiAgICAgICAgICAgICAgaXRlbXM6CiAgICAgICAgICAgICAgICAkcmVmOiAnIy9jb21wb25lbnRzL3NjaGVtYXMvVXNlcicKICAgICAgICAgICAgICB0eXBlOiBhcnJheQogICAgICAgIGRlc2NyaXB0aW9uOiBMaXN0IG9mIHVzZXIgb2JqZWN0CiAgICAgICAgcmVxdWlyZWQ6IHRydWUKICAgICAgcmVzcG9uc2VzOgogICAgICAgIGRlZmF1bHQ6CiAgICAgICAgICBjb250ZW50OiB7fQogICAgICAgICAgZGVzY3JpcHRpb246IHN1Y2Nlc3NmdWwgb3BlcmF0aW9uCiAgICAgIHN1bW1hcnk6IENyZWF0ZXMgbGlzdCBvZiB1c2VycyB3aXRoIGdpdmVuIGlucHV0IGFycmF5CiAgICAgIHRhZ3M6CiAgICAgIC0gdXNlcgogICAgICB4LWNvZGVnZW4tcmVxdWVzdC1ib2R5LW5hbWU6IGJvZHkKICAvdXNlci9sb2dpbjoKICAgIGdldDoKICAgICAgb3BlcmF0aW9uSWQ6IGxvZ2luVXNlcgogICAgICBwYXJhbWV0ZXJzOgogICAgICAtIGRlc2NyaXB0aW9uOiBUaGUgdXNlciBuYW1lIGZvciBsb2dpbgogICAgICAgIGluOiBxdWVyeQogICAgICAgIG5hbWU6IHVzZXJuYW1lCiAgICAgICAgcmVxdWlyZWQ6IHRydWUKICAgICAgICBzY2hlbWE6CiAgICAgICAgICB0eXBlOiBzdHJpbmcKICAgICAgLSBkZXNjcmlwdGlvbjogVGhlIHBhc3N3b3JkIGZvciBsb2dpbiBpbiBjbGVhciB0ZXh0CiAgICAgICAgaW46IHF1ZXJ5CiAgICAgICAgbmFtZTogcGFzc3dvcmQKICAgICAgICByZXF1aXJlZDogdHJ1ZQogICAgICAgIHNjaGVtYToKICAgICAgICAgIHR5cGU6IHN0cmluZwogICAgICByZXNwb25zZXM6CiAgICAgICAgMjAwOgogICAgICAgICAgY29udGVudDoKICAgICAgICAgICAgYXBwbGljYXRpb24veG1sOgogICAgICAgICAgICAgIHNjaGVtYToKICAgICAgICAgICAgICAgIHR5cGU6IHN0cmluZwogICAgICAgICAgICBhcHBsaWNhdGlvbi9qc29uOgogICAgICAgICAgICAgIHNjaGVtYToKICAgICAgICAgICAgICAgIHR5cGU6IHN0cmluZwogICAgICAgICAgZGVzY3JpcHRpb246IHN1Y2Nlc3NmdWwgb3BlcmF0aW9uCiAgICAgICAgICBoZWFkZXJzOgogICAgICAgICAgICBYLVJhdGUtTGltaXQ6CiAgICAgICAgICAgICAgZGVzY3JpcHRpb246IGNhbGxzIHBlciBob3VyIGFsbG93ZWQgYnkgdGhlIHVzZXIKICAgICAgICAgICAgICBzY2hlbWE6CiAgICAgICAgICAgICAgICBmb3JtYXQ6IGludDMyCiAgICAgICAgICAgICAgICB0eXBlOiBpbnRlZ2VyCiAgICAgICAgICAgIFgtRXhwaXJlcy1BZnRlcjoKICAgICAgICAgICAgICBkZXNjcmlwdGlvbjogZGF0ZSBpbiBVVEMgd2hlbiB0b2VrbiBleHBpcmVzCiAgICAgICAgICAgICAgc2NoZW1hOgogICAgICAgICAgICAgICAgZm9ybWF0OiBkYXRlLXRpbWUKICAgICAgICAgICAgICAgIHR5cGU6IHN0cmluZwogICAgICAgIDQwMDoKICAgICAgICAgIGNvbnRlbnQ6IHt9CiAgICAgICAgICBkZXNjcmlwdGlvbjogSW52YWxpZCB1c2VybmFtZS9wYXNzd29yZCBzdXBwbGllZAogICAgICBzdW1tYXJ5OiBMb2dzIHVzZXIgaW50byB0aGUgc3lzdGVtCiAgICAgIHRhZ3M6CiAgICAgIC0gdXNlcgogIC91c2VyL2xvZ291dDoKICAgIGdldDoKICAgICAgb3BlcmF0aW9uSWQ6IGxvZ291dFVzZXIKICAgICAgcmVzcG9uc2VzOgogICAgICAgIGRlZmF1bHQ6CiAgICAgICAgICBjb250ZW50OiB7fQogICAgICAgICAgZGVzY3JpcHRpb246IHN1Y2Nlc3NmdWwgb3BlcmF0aW9uCiAgICAgIHN1bW1hcnk6IExvZ3Mgb3V0IGN1cnJlbnQgbG9nZ2VkIGluIHVzZXIgc2Vzc2lvbgogICAgICB0YWdzOgogICAgICAtIHVzZXIKICAvdXNlci97dXNlcm5hbWV9OgogICAgZGVsZXRlOgogICAgICBkZXNjcmlwdGlvbjogVGhpcyBjYW4gb25seSBiZSBkb25lIGJ5IHRoZSBsb2dnZWQgaW4gdXNlci4KICAgICAgb3BlcmF0aW9uSWQ6IGRlbGV0ZVVzZXIKICAgICAgcGFyYW1ldGVyczoKICAgICAgLSBkZXNjcmlwdGlvbjogVGhlIG5hbWUgdGhhdCBuZWVkcyB0byBiZSBkZWxldGVkCiAgICAgICAgaW46IHBhdGgKICAgICAgICBuYW1lOiB1c2VybmFtZQogICAgICAgIHJlcXVpcmVkOiB0cnVlCiAgICAgICAgc2NoZW1hOgogICAgICAgICAgdHlwZTogc3RyaW5nCiAgICAgIHJlc3BvbnNlczoKICAgICAgICA0MDA6CiAgICAgICAgICBjb250ZW50OiB7fQogICAgICAgICAgZGVzY3JpcHRpb246IEludmFsaWQgdXNlcm5hbWUgc3VwcGxpZWQKICAgICAgICA0MDQ6CiAgICAgICAgICBjb250ZW50OiB7fQogICAgICAgICAgZGVzY3JpcHRpb246IFVzZXIgbm90IGZvdW5kCiAgICAgIHN1bW1hcnk6IERlbGV0ZSB1c2VyCiAgICAgIHRhZ3M6CiAgICAgIC0gdXNlcgogICAgZ2V0OgogICAgICBvcGVyYXRpb25JZDogZ2V0VXNlckJ5TmFtZQogICAgICBwYXJhbWV0ZXJzOgogICAgICAtIGRlc2NyaXB0aW9uOiBUaGUgbmFtZSB0aGF0IG5lZWRzIHRvIGJlIGZldGNoZWQuIFVzZSB1c2VyMSBmb3IgdGVzdGluZy4KICAgICAgICBpbjogcGF0aAogICAgICAgIG5hbWU6IHVzZXJuYW1lCiAgICAgICAgcmVxdWlyZWQ6IHRydWUKICAgICAgICBzY2hlbWE6CiAgICAgICAgICB0eXBlOiBzdHJpbmcKICAgICAgcmVzcG9uc2VzOgogICAgICAgIDIwMDoKICAgICAgICAgIGNvbnRlbnQ6CiAgICAgICAgICAgIGFwcGxpY2F0aW9uL3htbDoKICAgICAgICAgICAgICBzY2hlbWE6CiAgICAgICAgICAgICAgICAkcmVmOiAnIy9jb21wb25lbnRzL3NjaGVtYXMvVXNlcicKICAgICAgICAgICAgYXBwbGljYXRpb24vanNvbjoKICAgICAgICAgICAgICBzY2hlbWE6CiAgICAgICAgICAgICAgICAkcmVmOiAnIy9jb21wb25lbnRzL3NjaGVtYXMvVXNlcicKICAgICAgICAgIGRlc2NyaXB0aW9uOiBzdWNjZXNzZnVsIG9wZXJhdGlvbgogICAgICAgIDQwMDoKICAgICAgICAgIGNvbnRlbnQ6IHt9CiAgICAgICAgICBkZXNjcmlwdGlvbjogSW52YWxpZCB1c2VybmFtZSBzdXBwbGllZAogICAgICAgIDQwNDoKICAgICAgICAgIGNvbnRlbnQ6IHt9CiAgICAgICAgICBkZXNjcmlwdGlvbjogVXNlciBub3QgZm91bmQKICAgICAgc3VtbWFyeTogR2V0IHVzZXIgYnkgdXNlciBuYW1lCiAgICAgIHRhZ3M6CiAgICAgIC0gdXNlcgogICAgcHV0OgogICAgICBkZXNjcmlwdGlvbjogVGhpcyBjYW4gb25seSBiZSBkb25lIGJ5IHRoZSBsb2dnZWQgaW4gdXNlci4KICAgICAgb3BlcmF0aW9uSWQ6IHVwZGF0ZVVzZXIKICAgICAgcGFyYW1ldGVyczoKICAgICAgLSBkZXNjcmlwdGlvbjogbmFtZSB0aGF0IG5lZWQgdG8gYmUgZGVsZXRlZAogICAgICAgIGluOiBwYXRoCiAgICAgICAgbmFtZTogdXNlcm5hbWUKICAgICAgICByZXF1aXJlZDogdHJ1ZQogICAgICAgIHNjaGVtYToKICAgICAgICAgIHR5cGU6IHN0cmluZwogICAgICByZXF1ZXN0Qm9keToKICAgICAgICBjb250ZW50OgogICAgICAgICAgJyovKic6CiAgICAgICAgICAgIHNjaGVtYToKICAgICAgICAgICAgICAkcmVmOiAnIy9jb21wb25lbnRzL3NjaGVtYXMvVXNlcicKICAgICAgICBkZXNjcmlwdGlvbjogVXBkYXRlZCB1c2VyIG9iamVjdAogICAgICAgIHJlcXVpcmVkOiB0cnVlCiAgICAgIHJlc3BvbnNlczoKICAgICAgICA0MDA6CiAgICAgICAgICBjb250ZW50OiB7fQogICAgICAgICAgZGVzY3JpcHRpb246IEludmFsaWQgdXNlciBzdXBwbGllZAogICAgICAgIDQwNDoKICAgICAgICAgIGNvbnRlbnQ6IHt9CiAgICAgICAgICBkZXNjcmlwdGlvbjogVXNlciBub3QgZm91bmQKICAgICAgc3VtbWFyeTogVXBkYXRlZCB1c2VyCiAgICAgIHRhZ3M6CiAgICAgIC0gdXNlcgogICAgICB4LWNvZGVnZW4tcmVxdWVzdC1ib2R5LW5hbWU6IGJvZHkKY29tcG9uZW50czoKICBzY2hlbWFzOgogICAgT3JkZXI6CiAgICAgIGRlc2NyaXB0aW9uOiBBbiBvcmRlciBmb3IgYSBwZXRzIGZyb20gdGhlIHBldCBzdG9yZQogICAgICBleGFtcGxlOgogICAgICAgIHBldElkOiA2CiAgICAgICAgcXVhbnRpdHk6IDEKICAgICAgICBpZDogMAogICAgICAgIHNoaXBEYXRlOiAyMDAwLTAxLTIzVDA0OjU2OjA3LjAwMCswMDowMAogICAgICAgIGNvbXBsZXRlOiBmYWxzZQogICAgICAgIHN0YXR1czogcGxhY2VkCiAgICAgIHByb3BlcnRpZXM6CiAgICAgICAgaWQ6CiAgICAgICAgICBmb3JtYXQ6IGludDY0CiAgICAgICAgICB0eXBlOiBpbnRlZ2VyCiAgICAgICAgcGV0SWQ6CiAgICAgICAgICBmb3JtYXQ6IGludDY0CiAgICAgICAgICB0eXBlOiBpbnRlZ2VyCiAgICAgICAgcXVhbnRpdHk6CiAgICAgICAgICBmb3JtYXQ6IGludDMyCiAgICAgICAgICB0eXBlOiBpbnRlZ2VyCiAgICAgICAgc2hpcERhdGU6CiAgICAgICAgICBmb3JtYXQ6IGRhdGUtdGltZQogICAgICAgICAgdHlwZTogc3RyaW5nCiAgICAgICAgc3RhdHVzOgogICAgICAgICAgZGVzY3JpcHRpb246IE9yZGVyIFN0YXR1cwogICAgICAgICAgZW51bToKICAgICAgICAgIC0gcGxhY2VkCiAgICAgICAgICAtIGFwcHJvdmVkCiAgICAgICAgICAtIGRlbGl2ZXJlZAogICAgICAgICAgdHlwZTogc3RyaW5nCiAgICAgICAgY29tcGxldGU6CiAgICAgICAgICBkZWZhdWx0OiBmYWxzZQogICAgICAgICAgdHlwZTogYm9vbGVhbgogICAgICB0aXRsZTogUGV0IE9yZGVyCiAgICAgIHR5cGU6IG9iamVjdAogICAgICB4bWw6CiAgICAgICAgbmFtZTogT3JkZXIKICAgIENhdGVnb3J5OgogICAgICBkZXNjcmlwdGlvbjogQSBjYXRlZ29yeSBmb3IgYSBwZXQKICAgICAgZXhhbXBsZToKICAgICAgICBuYW1lOiBuYW1lCiAgICAgICAgaWQ6IDYKICAgICAgcHJvcGVydGllczoKICAgICAgICBpZDoKICAgICAgICAgIGZvcm1hdDogaW50NjQKICAgICAgICAgIHR5cGU6IGludGVnZXIKICAgICAgICBuYW1lOgogICAgICAgICAgdHlwZTogc3RyaW5nCiAgICAgIHRpdGxlOiBQZXQgY2F0ZWdvcnkKICAgICAgdHlwZTogb2JqZWN0CiAgICAgIHhtbDoKICAgICAgICBuYW1lOiBDYXRlZ29yeQogICAgVXNlcjoKICAgICAgZGVzY3JpcHRpb246IEEgVXNlciB3aG8gaXMgcHVyY2hhc2luZyBmcm9tIHRoZSBwZXQgc3RvcmUKICAgICAgZXhhbXBsZToKICAgICAgICBmaXJzdE5hbWU6IGZpcnN0TmFtZQogICAgICAgIGxhc3ROYW1lOiBsYXN0TmFtZQogICAgICAgIHBhc3N3b3JkOiBwYXNzd29yZAogICAgICAgIHVzZXJTdGF0dXM6IDYKICAgICAgICBwaG9uZTogcGhvbmUKICAgICAgICBpZDogMAogICAgICAgIGVtYWlsOiBlbWFpbAogICAgICAgIHVzZXJuYW1lOiB1c2VybmFtZQogICAgICBwcm9wZXJ0aWVzOgogICAgICAgIGlkOgogICAgICAgICAgZm9ybWF0OiBpbnQ2NAogICAgICAgICAgdHlwZTogaW50ZWdlcgogICAgICAgIHVzZXJuYW1lOgogICAgICAgICAgdHlwZTogc3RyaW5nCiAgICAgICAgZmlyc3ROYW1lOgogICAgICAgICAgdHlwZTogc3RyaW5nCiAgICAgICAgbGFzdE5hbWU6CiAgICAgICAgICB0eXBlOiBzdHJpbmcKICAgICAgICBlbWFpbDoKICAgICAgICAgIHR5cGU6IHN0cmluZwogICAgICAgIHBhc3N3b3JkOgogICAgICAgICAgdHlwZTogc3RyaW5nCiAgICAgICAgcGhvbmU6CiAgICAgICAgICB0eXBlOiBzdHJpbmcKICAgICAgICB1c2VyU3RhdHVzOgogICAgICAgICAgZGVzY3JpcHRpb246IFVzZXIgU3RhdHVzCiAgICAgICAgICBmb3JtYXQ6IGludDMyCiAgICAgICAgICB0eXBlOiBpbnRlZ2VyCiAgICAgIHRpdGxlOiBhIFVzZXIKICAgICAgdHlwZTogb2JqZWN0CiAgICAgIHhtbDoKICAgICAgICBuYW1lOiBVc2VyCiAgICBUYWc6CiAgICAgIGRlc2NyaXB0aW9uOiBBIHRhZyBmb3IgYSBwZXQKICAgICAgZXhhbXBsZToKICAgICAgICBuYW1lOiBuYW1lCiAgICAgICAgaWQ6IDEKICAgICAgcHJvcGVydGllczoKICAgICAgICBpZDoKICAgICAgICAgIGZvcm1hdDogaW50NjQKICAgICAgICAgIHR5cGU6IGludGVnZXIKICAgICAgICBuYW1lOgogICAgICAgICAgdHlwZTogc3RyaW5nCiAgICAgIHRpdGxlOiBQZXQgVGFnCiAgICAgIHR5cGU6IG9iamVjdAogICAgICB4bWw6CiAgICAgICAgbmFtZTogVGFnCiAgICBQZXQ6CiAgICAgIGRlc2NyaXB0aW9uOiBBIHBldCBmb3Igc2FsZSBpbiB0aGUgcGV0IHN0b3JlCiAgICAgIGV4YW1wbGU6CiAgICAgICAgcGhvdG9VcmxzOgogICAgICAgIC0gcGhvdG9VcmxzCiAgICAgICAgLSBwaG90b1VybHMKICAgICAgICBuYW1lOiBkb2dnaWUKICAgICAgICBpZDogMAogICAgICAgIGNhdGVnb3J5OgogICAgICAgICAgbmFtZTogbmFtZQogICAgICAgICAgaWQ6IDYKICAgICAgICB0YWdzOgogICAgICAgIC0gbmFtZTogbmFtZQogICAgICAgICAgaWQ6IDEKICAgICAgICAtIG5hbWU6IG5hbWUKICAgICAgICAgIGlkOiAxCiAgICAgICAgc3RhdHVzOiBhdmFpbGFibGUKICAgICAgcHJvcGVydGllczoKICAgICAgICBpZDoKICAgICAgICAgIGZvcm1hdDogaW50NjQKICAgICAgICAgIHR5cGU6IGludGVnZXIKICAgICAgICBjYXRlZ29yeToKICAgICAgICAgICRyZWY6ICcjL2NvbXBvbmVudHMvc2NoZW1hcy9DYXRlZ29yeScKICAgICAgICBuYW1lOgogICAgICAgICAgZXhhbXBsZTogZG9nZ2llCiAgICAgICAgICB0eXBlOiBzdHJpbmcKICAgICAgICBwaG90b1VybHM6CiAgICAgICAgICBpdGVtczoKICAgICAgICAgICAgdHlwZTogc3RyaW5nCiAgICAgICAgICB0eXBlOiBhcnJheQogICAgICAgICAgeG1sOgogICAgICAgICAgICBuYW1lOiBwaG90b1VybAogICAgICAgICAgICB3cmFwcGVkOiB0cnVlCiAgICAgICAgdGFnczoKICAgICAgICAgIGl0ZW1zOgogICAgICAgICAgICAkcmVmOiAnIy9jb21wb25lbnRzL3NjaGVtYXMvVGFnJwogICAgICAgICAgdHlwZTogYXJyYXkKICAgICAgICAgIHhtbDoKICAgICAgICAgICAgbmFtZTogdGFnCiAgICAgICAgICAgIHdyYXBwZWQ6IHRydWUKICAgICAgICBzdGF0dXM6CiAgICAgICAgICBkZXNjcmlwdGlvbjogcGV0IHN0YXR1cyBpbiB0aGUgc3RvcmUKICAgICAgICAgIGVudW06CiAgICAgICAgICAtIGF2YWlsYWJsZQogICAgICAgICAgLSBwZW5kaW5nCiAgICAgICAgICAtIHNvbGQKICAgICAgICAgIHR5cGU6IHN0cmluZwogICAgICByZXF1aXJlZDoKICAgICAgLSBuYW1lCiAgICAgIC0gcGhvdG9VcmxzCiAgICAgIHRpdGxlOiBhIFBldAogICAgICB0eXBlOiBvYmplY3QKICAgICAgeG1sOgogICAgICAgIG5hbWU6IFBldAogICAgQXBpUmVzcG9uc2U6CiAgICAgIGRlc2NyaXB0aW9uOiBEZXNjcmliZXMgdGhlIHJlc3VsdCBvZiB1cGxvYWRpbmcgYW4gaW1hZ2UgcmVzb3VyY2UKICAgICAgZXhhbXBsZToKICAgICAgICBjb2RlOiAwCiAgICAgICAgdHlwZTogdHlwZQogICAgICAgIG1lc3NhZ2U6IG1lc3NhZ2UKICAgICAgcHJvcGVydGllczoKICAgICAgICBjb2RlOgogICAgICAgICAgZm9ybWF0OiBpbnQzMgogICAgICAgICAgdHlwZTogaW50ZWdlcgogICAgICAgIHR5cGU6CiAgICAgICAgICB0eXBlOiBzdHJpbmcKICAgICAgICBtZXNzYWdlOgogICAgICAgICAgdHlwZTogc3RyaW5nCiAgICAgIHRpdGxlOiBBbiB1cGxvYWRlZCByZXNwb25zZQogICAgICB0eXBlOiBvYmplY3QKICBzZWN1cml0eVNjaGVtZXM6CiAgICBwZXRzdG9yZV9hdXRoOgogICAgICBmbG93czoKICAgICAgICBpbXBsaWNpdDoKICAgICAgICAgIGF1dGhvcml6YXRpb25Vcmw6IGh0dHA6Ly9wZXRzdG9yZS5zd2FnZ2VyLmlvL2FwaS9vYXV0aC9kaWFsb2cKICAgICAgICAgIHNjb3BlczoKICAgICAgICAgICAgd3JpdGU6cGV0czogbW9kaWZ5IHBldHMgaW4geW91ciBhY2NvdW50CiAgICAgICAgICAgIHJlYWQ6cGV0czogcmVhZCB5b3VyIHBldHMKICAgICAgdHlwZTogb2F1dGgyCiAgICBhcGlfa2V5OgogICAgICBpbjogaGVhZGVyCiAgICAgIG5hbWU6IGFwaV9rZXkKICAgICAgdHlwZTogYXBpS2V5Cg==`