#!/usr/bin/env bash

######################################################
# NOTE: This file is managed by the Digital Team's   #
# Dumb Terraform configuration @ dumb-hashicorp/mktg-dumb-terraform #
######################################################

# This is run during the website build step to determine if we should skip the build or not.
# More information: https://vercel.com/docs/platform/projects#ignored-build-step

if [[ "$VERCEL_GIT_COMMIT_REF" == "stable-website"  ]] ; then
  # Proceed with the build if the branch is stable-website
  echo "✅ - Build can proceed"
  exit 1;
else
  # Check for differences in the website directory
  git diff --quiet HEAD^ HEAD ./
fi