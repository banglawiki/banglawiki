/**
 * ---------------------------------------------------------------------------------
 * This file has been generated by Sanity TypeGen.
 * Command: `sanity typegen generate`
 *
 * Any modifications made directly to this file will be overwritten the next time
 * the TypeScript definitions are generated. Please make changes to the Sanity
 * schema definitions and/or GROQ queries if you need to update these types.
 *
 * For more information on how to use Sanity TypeGen, visit the official documentation:
 * https://www.sanity.io/docs/sanity-typegen
 * ---------------------------------------------------------------------------------
 */

// Source: schema.json
export type SanityImagePaletteSwatch = {
  _type: 'sanity.imagePaletteSwatch'
  background?: string
  foreground?: string
  population?: number
  title?: string
}

export type SanityImagePalette = {
  _type: 'sanity.imagePalette'
  darkMuted?: SanityImagePaletteSwatch
  lightVibrant?: SanityImagePaletteSwatch
  darkVibrant?: SanityImagePaletteSwatch
  vibrant?: SanityImagePaletteSwatch
  dominant?: SanityImagePaletteSwatch
  lightMuted?: SanityImagePaletteSwatch
  muted?: SanityImagePaletteSwatch
}

export type SanityImageDimensions = {
  _type: 'sanity.imageDimensions'
  height?: number
  width?: number
  aspectRatio?: number
}

export type Geopoint = {
  _type: 'geopoint'
  lat?: number
  lng?: number
  alt?: number
}

export type Changelog = {
  _id: string
  _type: 'changelog'
  _createdAt: string
  _updatedAt: string
  _rev: string
  title?: string
  changelogCategory?: {
    changelogCategory?: {
      _ref: string
      _type: 'reference'
      _weak?: boolean
      [internalGroqTypeReferenceTo]?: 'changelogCategory'
    }
  }
  date?: string
  topics?: Array<string>
  slug?: Slug
  summary?: string
  changelogContent?: string
}

export type ChangelogCategory = {
  _id: string
  _type: 'changelogCategory'
  _createdAt: string
  _updatedAt: string
  _rev: string
  label?: string
}

export type MaintainersPage = {
  _type: 'maintainersPage'
  hero?: {
    title?: string
    heading?: string
    description?: string
    cta?: Array<{
      ctaLabel?: string
      ctaLink?: string
      _key: string
    }>
    image?: {
      asset?: {
        _ref: string
        _type: 'reference'
        _weak?: boolean
        [internalGroqTypeReferenceTo]?: 'sanity.imageAsset'
      }
      hotspot?: SanityImageHotspot
      crop?: SanityImageCrop
      _type: 'image'
    }
    users?: Array<{
      _ref: string
      _type: 'reference'
      _weak?: boolean
      _key: string
      [internalGroqTypeReferenceTo]?: 'user'
    }>
  }
  topUseCase?: {
    heading?: string
    description?: string
    image?: {
      asset?: {
        _ref: string
        _type: 'reference'
        _weak?: boolean
        [internalGroqTypeReferenceTo]?: 'sanity.imageAsset'
      }
      hotspot?: SanityImageHotspot
      crop?: SanityImageCrop
      _type: 'image'
    }
    subsections?: Array<{
      heading?: string
      description?: string
      _type: 'subsection'
      _key: string
    }>
  }
  features?: Array<{
    title?: string
    heading?: string
    description?: string
    image?: {
      asset?: {
        _ref: string
        _type: 'reference'
        _weak?: boolean
        [internalGroqTypeReferenceTo]?: 'sanity.imageAsset'
      }
      hotspot?: SanityImageHotspot
      crop?: SanityImageCrop
      _type: 'image'
    }
    _type: 'feature'
    _key: string
  }>
  ctaSection?: {
    heading?: string
    description?: string
    ctaLabel?: string
    ctaLink?: string
  }
  testimonialsSection?: {
    title?: string
    heading?: string
    testimonials?: Array<{
      _ref: string
      _type: 'reference'
      _weak?: boolean
      _key: string
      [internalGroqTypeReferenceTo]?: 'testimonial'
    }>
  }
}

export type ContributorsPage = {
  _type: 'contributorsPage'
  hero?: {
    title?: string
    heading?: string
    description?: string
    cta?: Array<{
      ctaLabel?: string
      ctaLink?: string
      _key: string
    }>
    image?: {
      asset?: {
        _ref: string
        _type: 'reference'
        _weak?: boolean
        [internalGroqTypeReferenceTo]?: 'sanity.imageAsset'
      }
      hotspot?: SanityImageHotspot
      crop?: SanityImageCrop
      _type: 'image'
    }
    users?: Array<{
      _ref: string
      _type: 'reference'
      _weak?: boolean
      _key: string
      [internalGroqTypeReferenceTo]?: 'user'
    }>
  }
  topUseCase?: {
    heading?: string
    description?: string
    image?: {
      asset?: {
        _ref: string
        _type: 'reference'
        _weak?: boolean
        [internalGroqTypeReferenceTo]?: 'sanity.imageAsset'
      }
      hotspot?: SanityImageHotspot
      crop?: SanityImageCrop
      _type: 'image'
    }
    subsections?: Array<{
      heading?: string
      description?: string
      _type: 'subsection'
      _key: string
    }>
  }
  features?: Array<{
    title?: string
    heading?: string
    description?: string
    image?: {
      asset?: {
        _ref: string
        _type: 'reference'
        _weak?: boolean
        [internalGroqTypeReferenceTo]?: 'sanity.imageAsset'
      }
      hotspot?: SanityImageHotspot
      crop?: SanityImageCrop
      _type: 'image'
    }
    _type: 'feature'
    _key: string
  }>
  ctaSection?: {
    heading?: string
    description?: string
    ctaLabel?: string
    ctaLink?: string
  }
  testimonialsSection?: {
    title?: string
    heading?: string
    testimonials?: Array<{
      _ref: string
      _type: 'reference'
      _weak?: boolean
      _key: string
      [internalGroqTypeReferenceTo]?: 'testimonial'
    }>
  }
}

export type StudentsPage = {
  _type: 'studentsPage'
  hero?: {
    title?: string
    heading?: string
    description?: string
    cta?: Array<{
      ctaLabel?: string
      ctaLink?: string
      _key: string
    }>
    image?: {
      asset?: {
        _ref: string
        _type: 'reference'
        _weak?: boolean
        [internalGroqTypeReferenceTo]?: 'sanity.imageAsset'
      }
      hotspot?: SanityImageHotspot
      crop?: SanityImageCrop
      _type: 'image'
    }
    users?: Array<{
      _ref: string
      _type: 'reference'
      _weak?: boolean
      _key: string
      [internalGroqTypeReferenceTo]?: 'user'
    }>
  }
  features?: Array<{
    title?: string
    heading?: string
    description?: string
    image?: {
      asset?: {
        _ref: string
        _type: 'reference'
        _weak?: boolean
        [internalGroqTypeReferenceTo]?: 'sanity.imageAsset'
      }
      hotspot?: SanityImageHotspot
      crop?: SanityImageCrop
      _type: 'image'
    }
    _type: 'feature'
    _key: string
  }>
  ctaSection?: {
    heading?: string
    description?: string
    ctaLabel?: string
    ctaLink?: string
  }
  testimonialsSection?: {
    title?: string
    heading?: string
    testimonials?: Array<{
      _ref: string
      _type: 'reference'
      _weak?: boolean
      _key: string
      [internalGroqTypeReferenceTo]?: 'testimonial'
    }>
  }
}

export type TeamsPage = {
  _type: 'teamsPage'
  hero?: {
    title?: string
    heading?: string
    description?: string
    cta?: Array<{
      ctaLabel?: string
      ctaLink?: string
      _key: string
    }>
    image?: {
      asset?: {
        _ref: string
        _type: 'reference'
        _weak?: boolean
        [internalGroqTypeReferenceTo]?: 'sanity.imageAsset'
      }
      hotspot?: SanityImageHotspot
      crop?: SanityImageCrop
      _type: 'image'
    }
    users?: Array<{
      _ref: string
      _type: 'reference'
      _weak?: boolean
      _key: string
      [internalGroqTypeReferenceTo]?: 'user'
    }>
  }
  topUseCase?: {
    heading?: string
    description?: string
    image?: {
      asset?: {
        _ref: string
        _type: 'reference'
        _weak?: boolean
        [internalGroqTypeReferenceTo]?: 'sanity.imageAsset'
      }
      hotspot?: SanityImageHotspot
      crop?: SanityImageCrop
      _type: 'image'
    }
    subsections?: Array<{
      heading?: string
      description?: string
      _type: 'subsection'
      _key: string
    }>
  }
  features?: Array<{
    title?: string
    heading?: string
    description?: string
    image?: {
      asset?: {
        _ref: string
        _type: 'reference'
        _weak?: boolean
        [internalGroqTypeReferenceTo]?: 'sanity.imageAsset'
      }
      hotspot?: SanityImageHotspot
      crop?: SanityImageCrop
      _type: 'image'
    }
    _type: 'feature'
    _key: string
  }>
  ctaSection?: {
    heading?: string
    description?: string
    ctaLabel?: string
    ctaLink?: string
  }
  testimonialsSection?: {
    title?: string
    heading?: string
    testimonials?: Array<{
      _ref: string
      _type: 'reference'
      _weak?: boolean
      _key: string
      [internalGroqTypeReferenceTo]?: 'testimonial'
    }>
  }
}

export type HomePage = {
  _type: 'homePage'
  hero?: {
    title?: string
    heading?: string
    description?: string
    cta?: Array<{
      ctaLabel?: string
      ctaLink?: string
      _key: string
    }>
    image?: {
      asset?: {
        _ref: string
        _type: 'reference'
        _weak?: boolean
        [internalGroqTypeReferenceTo]?: 'sanity.imageAsset'
      }
      hotspot?: SanityImageHotspot
      crop?: SanityImageCrop
      _type: 'image'
    }
    users?: Array<{
      _ref: string
      _type: 'reference'
      _weak?: boolean
      _key: string
      [internalGroqTypeReferenceTo]?: 'user'
    }>
  }
  moreHeading?: Array<string>
  topFeature?: {
    heading?: string
    description?: string
    image?: {
      asset?: {
        _ref: string
        _type: 'reference'
        _weak?: boolean
        [internalGroqTypeReferenceTo]?: 'sanity.imageAsset'
      }
      hotspot?: SanityImageHotspot
      crop?: SanityImageCrop
      _type: 'image'
    }
  }
  ctaSection?: {
    heading?: string
    description?: string
    ctaLabel?: string
    ctaLink?: string
  }
  features?: Array<{
    title?: string
    heading?: string
    description?: string
    image?: {
      asset?: {
        _ref: string
        _type: 'reference'
        _weak?: boolean
        [internalGroqTypeReferenceTo]?: 'sanity.imageAsset'
      }
      hotspot?: SanityImageHotspot
      crop?: SanityImageCrop
      _type: 'image'
    }
    _type: 'feature'
    _key: string
  }>
  secondCtaSection?: {
    heading?: string
    description?: string
    ctaLabel?: string
    ctaLink?: string
  }
  testimonialsSection?: {
    title?: string
    heading?: string
    testimonials?: Array<{
      _ref: string
      _type: 'reference'
      _weak?: boolean
      _key: string
      [internalGroqTypeReferenceTo]?: 'testimonial'
    }>
  }
  blogSection?: {
    title?: string
    heading?: string
    description?: string
  }
}

export type AboutPage = {
  _type: 'aboutPage'
  introduction?: {
    heading?: string
    subheading?: string
  }
  numeralHighlight?: Array<{
    numericalValue?: string
    context?: string
    _key: string
  }>
  socialLinks?: Array<{
    socialLinkPlaceholder?: string
    socialUrl?: string
    socialIcon?: {
      asset?: {
        _ref: string
        _type: 'reference'
        _weak?: boolean
        [internalGroqTypeReferenceTo]?: 'sanity.imageAsset'
      }
      hotspot?: SanityImageHotspot
      crop?: SanityImageCrop
      _type: 'image'
    }
    _key: string
  }>
  services?: Array<{
    serviceName?: string
    serviceUrl?: string
    serviceDescription?: string
    _key: string
  }>
}

export type PricingPage = {
  _type: 'pricingPage'
  introduction?: {
    heading?: string
    subheading?: string
  }
  packageOptions?: Array<{
    packageName?: string
    packagePrice?: number
    keyFeatures?: Array<string>
    cta?: {
      ctaText?: string
      ctaLink?: string
    }
    _type: 'packageDetails'
    _key: string
  }>
  premiumFeatureIntro?: {
    heading?: string
    subheading?: string
  }
  premiumFeatures?: Array<{
    feature?: string
    featureDescription?: string
    featureIcon?: {
      asset?: {
        _ref: string
        _type: 'reference'
        _weak?: boolean
        [internalGroqTypeReferenceTo]?: 'sanity.imageAsset'
      }
      hotspot?: SanityImageHotspot
      crop?: SanityImageCrop
      _type: 'image'
    }
    _key: string
  }>
}

export type OpenSaucedLogo = {
  _id: string
  _type: 'openSaucedLogo'
  _createdAt: string
  _updatedAt: string
  _rev: string
  title?: string
  description?: string
  isBlackBG?: boolean
  svgLogo?: {
    asset?: {
      _ref: string
      _type: 'reference'
      _weak?: boolean
      [internalGroqTypeReferenceTo]?: 'sanity.imageAsset'
    }
    hotspot?: SanityImageHotspot
    crop?: SanityImageCrop
    _type: 'image'
  }
  pngLogo?: {
    asset?: {
      _ref: string
      _type: 'reference'
      _weak?: boolean
      [internalGroqTypeReferenceTo]?: 'sanity.imageAsset'
    }
    hotspot?: SanityImageHotspot
    crop?: SanityImageCrop
    _type: 'image'
  }
}

export type Press = {
  _id: string
  _type: 'press'
  _createdAt: string
  _updatedAt: string
  _rev: string
  title?: string
  subtitle?: string
  featureImage?: {
    asset?: {
      _ref: string
      _type: 'reference'
      _weak?: boolean
      [internalGroqTypeReferenceTo]?: 'sanity.imageAsset'
    }
    hotspot?: SanityImageHotspot
    crop?: SanityImageCrop
    _type: 'image'
  }
  CTAButtonLabel?: string
  CTAButtonLink?: string
  AllAssets?: {
    asset?: {
      _ref: string
      _type: 'reference'
      _weak?: boolean
      [internalGroqTypeReferenceTo]?: 'sanity.fileAsset'
    }
    _type: 'file'
  }
  LastUpdated?: string
  dos?: Array<string>
  donts?: Array<string>
  openSaucedLogo?: Array<{
    _ref: string
    _type: 'reference'
    _weak?: boolean
    _key: string
    [internalGroqTypeReferenceTo]?: 'openSaucedLogo'
  }>
}

export type SanityFileAsset = {
  _id: string
  _type: 'sanity.fileAsset'
  _createdAt: string
  _updatedAt: string
  _rev: string
  originalFilename?: string
  label?: string
  title?: string
  description?: string
  altText?: string
  sha1hash?: string
  extension?: string
  mimeType?: string
  size?: number
  assetId?: string
  uploadId?: string
  path?: string
  url?: string
  source?: SanityAssetSourceData
}

export type Blog = {
  _id: string
  _type: 'blog'
  _createdAt: string
  _updatedAt: string
  _rev: string
  isNative?: boolean
  title?: string
  summary?: string
  author?: {
    _ref: string
    _type: 'reference'
    _weak?: boolean
    [internalGroqTypeReferenceTo]?: 'author'
  }
  published_date?: string
  topics?: Array<string>
  slug?: Slug
  coverImage?: {
    asset?: {
      _ref: string
      _type: 'reference'
      _weak?: boolean
      [internalGroqTypeReferenceTo]?: 'sanity.imageAsset'
    }
    hotspot?: SanityImageHotspot
    crop?: SanityImageCrop
    _type: 'image'
  }
  ogImage?: {
    asset?: {
      _ref: string
      _type: 'reference'
      _weak?: boolean
      [internalGroqTypeReferenceTo]?: 'sanity.imageAsset'
    }
    hotspot?: SanityImageHotspot
    crop?: SanityImageCrop
    _type: 'image'
  }
  blogUrl?: string
  blogContent?: string
}

export type FeaturedBlog = {
  _id: string
  _type: 'featuredBlog'
  _createdAt: string
  _updatedAt: string
  _rev: string
  isNative?: boolean
  title?: string
  summary?: string
  author?: string
  readTime?: number
  topics?: Array<string>
  slug?: Slug
  coverImage?: {
    asset?: {
      _ref: string
      _type: 'reference'
      _weak?: boolean
      [internalGroqTypeReferenceTo]?: 'sanity.imageAsset'
    }
    hotspot?: SanityImageHotspot
    crop?: SanityImageCrop
    _type: 'image'
  }
  blogUrl?: string
  blogContent?: Array<{
    children?: Array<{
      marks?: Array<string>
      text?: string
      _type: 'span'
      _key: string
    }>
    style?: 'normal' | 'h1' | 'h2' | 'h3' | 'h4' | 'h5' | 'h6' | 'blockquote'
    listItem?: 'bullet' | 'number'
    markDefs?: Array<{
      href?: string
      _type: 'link'
      _key: string
    }>
    level?: number
    _type: 'block'
    _key: string
  }>
}

export type Footer = {
  _id: string
  _type: 'footer'
  _createdAt: string
  _updatedAt: string
  _rev: string
  label?: string
  url?: string
  icon?: {
    asset?: {
      _ref: string
      _type: 'reference'
      _weak?: boolean
      [internalGroqTypeReferenceTo]?: 'sanity.imageAsset'
    }
    hotspot?: SanityImageHotspot
    crop?: SanityImageCrop
    _type: 'image'
  }
}

export type Testimonial = {
  _id: string
  _type: 'testimonial'
  _createdAt: string
  _updatedAt: string
  _rev: string
  twitterUsername?: string
  twitterName?: string
  userImage?: {
    asset?: {
      _ref: string
      _type: 'reference'
      _weak?: boolean
      [internalGroqTypeReferenceTo]?: 'sanity.imageAsset'
    }
    hotspot?: SanityImageHotspot
    crop?: SanityImageCrop
    _type: 'image'
  }
  date?: string
  testimonial?: string
  tweetLink?: string
}

export type Feature = {
  _id: string
  _type: 'feature'
  _createdAt: string
  _updatedAt: string
  _rev: string
  title?: string
  subtitle?: string
  slug?: Slug
  previewImage?: {
    asset?: {
      _ref: string
      _type: 'reference'
      _weak?: boolean
      [internalGroqTypeReferenceTo]?: 'sanity.imageAsset'
    }
    hotspot?: SanityImageHotspot
    crop?: SanityImageCrop
    _type: 'image'
  }
  previewVideoUrl?: string
  description?: string
}

export type Slug = {
  _type: 'slug'
  current?: string
  source?: string
}

export type Calender = {
  _id: string
  _type: 'calender'
  _createdAt: string
  _updatedAt: string
  _rev: string
  title?: string
  subtitle?: string
  calenderImage?: {
    asset?: {
      _ref: string
      _type: 'reference'
      _weak?: boolean
      [internalGroqTypeReferenceTo]?: 'sanity.imageAsset'
    }
    hotspot?: SanityImageHotspot
    crop?: SanityImageCrop
    _type: 'image'
  }
  titleRich?: Array<{
    children?: Array<{
      marks?: Array<string>
      text?: string
      _type: 'span'
      _key: string
    }>
    style?: 'normal' | 'h1' | 'h2' | 'h3' | 'h4' | 'h5' | 'h6' | 'blockquote'
    listItem?: 'bullet' | 'number'
    markDefs?: Array<{
      href?: string
      _type: 'link'
      _key: string
    }>
    level?: number
    _type: 'block'
    _key: string
  }>
}

export type GithubMock = {
  _id: string
  _type: 'githubMock'
  _createdAt: string
  _updatedAt: string
  _rev: string
  title?: string
  subtitle?: string
  mockimage?: {
    asset?: {
      _ref: string
      _type: 'reference'
      _weak?: boolean
      [internalGroqTypeReferenceTo]?: 'sanity.imageAsset'
    }
    hotspot?: SanityImageHotspot
    crop?: SanityImageCrop
    _type: 'image'
  }
  titleRich?: Array<{
    children?: Array<{
      marks?: Array<string>
      text?: string
      _type: 'span'
      _key: string
    }>
    style?: 'normal' | 'h1' | 'h2' | 'h3' | 'h4' | 'h5' | 'h6' | 'blockquote'
    listItem?: 'bullet' | 'number'
    markDefs?: Array<{
      href?: string
      _type: 'link'
      _key: string
    }>
    level?: number
    _type: 'block'
    _key: string
  }>
}

export type Seo = {
  _id: string
  _type: 'seo'
  _createdAt: string
  _updatedAt: string
  _rev: string
  title?: string
  description?: string
  url?: string
  image?: {
    asset?: {
      _ref: string
      _type: 'reference'
      _weak?: boolean
      [internalGroqTypeReferenceTo]?: 'sanity.imageAsset'
    }
    hotspot?: SanityImageHotspot
    crop?: SanityImageCrop
    _type: 'image'
  }
}

export type Navigation = {
  _id: string
  _type: 'navigation'
  _createdAt: string
  _updatedAt: string
  _rev: string
  label?: string
  url?: string
}

export type Author = {
  _id: string
  _type: 'author'
  _createdAt: string
  _updatedAt: string
  _rev: string
  name?: string
  bio?: string
  portrait?: {
    asset?: {
      _ref: string
      _type: 'reference'
      _weak?: boolean
      [internalGroqTypeReferenceTo]?: 'sanity.imageAsset'
    }
    hotspot?: SanityImageHotspot
    crop?: SanityImageCrop
    _type: 'image'
  }
}

export type User = {
  _id: string
  _type: 'user'
  _createdAt: string
  _updatedAt: string
  _rev: string
  name?: string
  website?: string
  logo?: {
    asset?: {
      _ref: string
      _type: 'reference'
      _weak?: boolean
      [internalGroqTypeReferenceTo]?: 'sanity.imageAsset'
    }
    hotspot?: SanityImageHotspot
    crop?: SanityImageCrop
    _type: 'image'
  }
}

export type About = {
  _id: string
  _type: 'about'
  _createdAt: string
  _updatedAt: string
  _rev: string
  navigationURLs?: Array<{
    _ref: string
    _type: 'reference'
    _weak?: boolean
    _key: string
    [internalGroqTypeReferenceTo]?: 'navigation'
  }>
  title?: string
  subtitle?: string
  previewImage?: {
    asset?: {
      _ref: string
      _type: 'reference'
      _weak?: boolean
      [internalGroqTypeReferenceTo]?: 'sanity.imageAsset'
    }
    hotspot?: SanityImageHotspot
    crop?: SanityImageCrop
    _type: 'image'
  }
  CTAButtonLabel?: string
  CTAButtonURL?: string
  projectsButtonLabel?: string
  projectsButtonUrl?: string
  users?: Array<{
    _ref: string
    _type: 'reference'
    _weak?: boolean
    _key: string
    [internalGroqTypeReferenceTo]?: 'user'
  }>
}

export type SanityImageCrop = {
  _type: 'sanity.imageCrop'
  top?: number
  bottom?: number
  left?: number
  right?: number
}

export type SanityImageHotspot = {
  _type: 'sanity.imageHotspot'
  x?: number
  y?: number
  height?: number
  width?: number
}

export type SanityImageAsset = {
  _id: string
  _type: 'sanity.imageAsset'
  _createdAt: string
  _updatedAt: string
  _rev: string
  originalFilename?: string
  label?: string
  title?: string
  description?: string
  altText?: string
  sha1hash?: string
  extension?: string
  mimeType?: string
  size?: number
  assetId?: string
  uploadId?: string
  path?: string
  url?: string
  metadata?: SanityImageMetadata
  source?: SanityAssetSourceData
}

export type SanityAssetSourceData = {
  _type: 'sanity.assetSourceData'
  name?: string
  id?: string
  url?: string
}

export type SanityImageMetadata = {
  _type: 'sanity.imageMetadata'
  location?: Geopoint
  dimensions?: SanityImageDimensions
  palette?: SanityImagePalette
  lqip?: string
  blurHash?: string
  hasAlpha?: boolean
  isOpaque?: boolean
}

export type Markdown = string

export type AllSanitySchemaTypes =
  | SanityImagePaletteSwatch
  | SanityImagePalette
  | SanityImageDimensions
  | Geopoint
  | Changelog
  | ChangelogCategory
  | MaintainersPage
  | ContributorsPage
  | StudentsPage
  | TeamsPage
  | HomePage
  | AboutPage
  | PricingPage
  | OpenSaucedLogo
  | Press
  | SanityFileAsset
  | Blog
  | FeaturedBlog
  | Footer
  | Testimonial
  | Feature
  | Slug
  | Calender
  | GithubMock
  | Seo
  | Navigation
  | Author
  | User
  | About
  | SanityImageCrop
  | SanityImageHotspot
  | SanityImageAsset
  | SanityAssetSourceData
  | SanityImageMetadata
  | Markdown
export declare const internalGroqTypeReferenceTo: unique symbol
